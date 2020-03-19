package api

import (
	"admin-server/config"
	"encoding/json"
	dbclient "admin-server/service/db/client"
	"admin-server/service/db/orm"

	"admin-server/mq"
	"admin-server/store/oss"
	"admin-server/util"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 文件上传请求
func UploadHandler(c *gin.Context) {
	res := util.InitRpc()
	defer util.DeferRpc(&res, c)

	// 接受文件流及存储到本地目录
	file, head, err := c.Request.FormFile("file")
	if err != nil {
		msg := "Failed to get data, err: " + err.Error()
		log.Print(msg)
		res.Code = -1
		res.Msg = msg
		return
	}
	defer file.Close()

	fileMeta := dbclient.FileMeta{
		FileName: head.Filename,
		FileAddr: config.FilePath + head.Filename,
		UploadAt: time.Now().Format("2020-01-01 00:00:00"),
	}

	newFile, err := os.Create(fileMeta.FileAddr)
	if err != nil {
		msg := "Failed to create file, err: " + err.Error()
		log.Print(msg)
		res.Code = -1
		res.Msg = msg
		return
	}
	defer newFile.Close()

	fileMeta.FileSize, err = io.Copy(newFile, file)
	if err != nil {
		msg := "Failed to save data into file, err: " + err.Error()
		log.Print(msg)
		res.Code = -1
		res.Msg = msg
		return
	}
	newFile.Seek(0, 0)

	// 上传到 oss
	fileMeta.FileSha1 = util.FileSha1(newFile)
	ossKey := fileMeta.FileName + "-" + fileMeta.FileSha1

	// kafka
	data := mq.TransferData{
		FileName:     fileMeta.FileName,
		FileHash:     fileMeta.FileSha1,
		CurLocation:  fileMeta.FileAddr,
		DestLocation: ossKey,
	}

	pubData, _ := json.Marshal(data)
	mqsuc := mq.Publist(pubData)

	if !mqsuc {
		msg := "上传文件发布失败: " + err.Error()
		log.Print(msg)
		res.Code = -1
		res.Msg = msg
		return
	}

	_, error := dbclient.OnFileUploadFinished(fileMeta)
	if error != "" {
		msg := "上传文件发布失败: " + err.Error()
		log.Print(msg)
		res.Code = -1
		res.Msg = msg
		return
	}

	username, _ := c.Request.Cookie("username")
	_, error = dbclient.OnUserFileUploadFinished(username.Value, fileMeta)
	if error == "" {
		res.Msg = "Upload Success, sha1=" + fileMeta.FileSha1
	} else {
		res.Code = -1
		res.Msg = "上传文件发布失败"
	}
}

// 获取文件元信息
func GetFileMetaHandler(c *gin.Context) {
	res := util.InitRpc()
	defer util.DeferRpc(&res, c)

	filehash := c.Request.FormValue("filehash")

	dbfileMeta, err := dbclient.GetFileMeta(filehash)
	if err != "" {
		res.Code = -1
		res.Msg = "FAILED: " + err
		return
	}
	fMeta := dbclient.TableFileToFileMeta(dbfileMeta.Data.(orm.TableFile))

	data, _ := json.Marshal(map[string]interface{}{
		"fileMeta": fMeta,
	})

	res.Data = data
}

// 批量查询文件原信息
func FileQueryHandler(c *gin.Context) {
	res := util.InitRpc()
	defer util.DeferRpc(&res, c)

	limitCount, _ := strconv.Atoi(c.Request.FormValue("limit"))
	username, _ := c.Request.Cookie("username")

	dbResp, err := dbclient.QueryUserFileMetas(username.Value, limitCount)
	if err != "" {
		msg := "FAILED, err: " + err
		log.Print(msg)
		res.Code = -1
		res.Msg = msg
		return
	}

	data, _ := json.Marshal(map[string]interface{}{
		"list": func ()interface{} {
			if dbResp.Data == nil {
				return []interface{}{}
			}
			return dbResp.Data
		}(),
	})

	res.Data = data
}

// 重命名文件
func FileMetaRenameHandler(c *gin.Context) {
	res := util.InitRpc()
	defer util.DeferRpc(&res, c)

	username, _ := c.Request.Cookie("username")
	opType := c.Request.FormValue("optype")
	fileSha1 := c.Request.FormValue("filehash")
	newFileName := c.Request.FormValue("filename")

	if opType != "0" {
		res.Code = -1
		res.Msg = "FAILED: opType"
		return
	}

	// 数据库重命名
	_, err := dbclient.RenameFileName(username.Value, fileSha1, newFileName)
	if err != "" {
		res.Code = -1
		res.Msg = "重命名失败: " + err
		return
	}
}

// 删除文件
func FileDeleteHandle(c *gin.Context) {
	res := util.InitRpc()
	defer util.DeferRpc(&res, c)

	username, _ := c.Request.Cookie("username")
	filehash := c.Request.FormValue("filehash")

	_, err := dbclient.DeleteUserFile(username.Value, filehash)
	if err != "" {
		msg := "删除文件失败: " + err
		log.Print(msg)
		res.Code = -1
		res.Msg = msg
		return
	}
}

// 尝试秒传
func TryFastUploadHandler(c *gin.Context) {
	res := util.InitRpc()
	defer util.DeferRpc(&res, c)

	username, _ := c.Request.Cookie("username")
	filehash := c.Request.FormValue("filehash")

	dbfileMeta, err := dbclient.GetFileMeta(filehash)
	if err != "" {
		res.Code = -1
		res.Msg = "FAILED: " + err
		return
	}

	fmeta := dbclient.TableFileToFileMeta(dbclient.ToTableFile(dbfileMeta.Data.(orm.TableFile)))
	_, err = dbclient.OnUserFileUploadFinished(username.Value, fmeta)

	if err == "" {
		res.Msg = "秒传成功"
		c.JSON(http.StatusOK, util.NewRespMsg(0, "秒传成功", nil))
		return
	} else {
		res.Code = -1
		res.Msg = "秒传失败, 请稍后重试"
		return
	}
}

// 下载文件
func DownloadUrlHandler(c *gin.Context) {
	res := util.InitRpc()
	defer util.DeferRpc(&res, c)

	filehash := c.Request.FormValue("filehash")

	// 从文件表查记录
	dbfileMeta, err := dbclient.GetFileMeta(filehash)
	log.Print(dbfileMeta)
	log.Print(err)
	if err != "" {
		res.Code = -1
		res.Msg = "Faild: " + err
	}

	fmeta := dbclient.TableFileToFileMeta(dbclient.ToTableFile(dbfileMeta.Data))

	key := fmeta.FileAddr
	url := oss.DownloadPresignedUrl(key)

	data, _ := json.Marshal(map[string]interface{}{
		"url": url,
	})
	res.Data = data
}
