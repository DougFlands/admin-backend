package api

import (
	"admin-server/config"
	"encoding/json"
	rPool "admin-server/cache/redis"
	dbclient "admin-server/service/db/client"

	"admin-server/util"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// 初始化信息
type MultiparUploadInfo struct {
	FileHash   string
	FileSize   int
	UploadID   string
	ChunkSize  int // 分块大小
	ChunkCount int // 分块数量
}

// 初始化
func InitialMultipartUploadHandler(c *gin.Context) {
	res := util.InitRpc()
	defer util.DeferRpc(&res, c)

	username := c.Request.FormValue("username")
	filehash := c.Request.FormValue("filehash")
	filesize, err := strconv.Atoi(c.Request.FormValue("filesize"))

	if err != nil {
		msg := "Params filesize invalid: " + err.Error()
		log.Print(msg)
		res.Code = -1
		res.Msg = msg
		return
	}

	// 建立链接
	rConn := rPool.RedisPool().Get()
	defer rConn.Close()

	size := 5 * 1024 * 1024 // 5M 分块
	upInfo := MultiparUploadInfo{
		FileHash:   filehash,
		FileSize:   filesize,
		UploadID:   username + fmt.Sprint(time.Now().UnixNano()),
		ChunkSize:  size,                                              // 5M
		ChunkCount: int(math.Ceil(float64(filesize) / float64(size))), // 分成多少块，向上取整
	}

	// 存到 redis 里
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "chunkcount", upInfo.ChunkCount)
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "filehash", upInfo.FileHash)
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "filesize", upInfo.FileSize)

	data, _ := json.Marshal(map[string]interface{}{
		"info": upInfo,
	})

	res.Data = data
}

// 上传文件分块
func UploadPartHandler(c *gin.Context) {
	res := util.InitRpc()
	defer util.DeferRpc(&res, c)

	upLoadID := c.Request.FormValue("uploadid")
	chunkIndex := c.Request.FormValue("index")

	rConn := rPool.RedisPool().Get()
	defer rConn.Close()
	fpath := config.FilePath + upLoadID + "/" + chunkIndex
	os.MkdirAll(path.Dir(fpath), 0744)
	fd, err := os.Create(fpath)
	if err != nil {
		msg := "Upload Part Failed: " + err.Error()
		log.Print(msg)
		res.Code = -1
		res.Msg = msg
		return
	}

	defer fd.Close()

	buf := make([]byte, 1024*1024) // 每次读 1M
	for {
		n, err := c.Request.Body.Read(buf)
		fd.Write(buf[:n])
		if err != nil {
			break
		}
	}

	rConn.Do("HSET", "MP_"+upLoadID, "chunkindex_"+chunkIndex, 1)
}

// 通知上传合并
func CompleteUploadHandler(c *gin.Context) {
	res := util.InitRpc()
	defer util.DeferRpc(&res, c)

	upLoadID := c.Request.FormValue("uploadid")
	chunkIndex := c.Request.FormValue("index")
	username := c.Request.FormValue("username")
	fileHash := c.Request.FormValue("filehash")
	fileSize := c.Request.FormValue("filesize")
	fileName := c.Request.FormValue("filename")

	rConn := rPool.RedisPool().Get()
	defer rConn.Close()

	data, err := redis.Values(rConn.Do("HGETALL", "MP_"+upLoadID))
	if err != nil {
		msg := "Complete Upload Failed: " + err.Error()
		log.Print(msg)
		res.Code = -1
		res.Msg = msg
		return
	}

	//  实际存了多少块
	totalCount := 0
	// 总共分了多少块
	chunkCount := 0

	// 返回第一个值为key 第二个为value
	for i := 0; i < len(data); i += 2 {
		k := string(data[i].([]byte))
		v := string(data[i+1].([]byte))
		if k == "chunkcount" {
			chunkCount, _ = strconv.Atoi(v)
		} else if strings.HasPrefix(k, "chunkindex_") && v == "1" {
			totalCount++
		}
	}
	fmt.Print(totalCount, chunkCount)

	if totalCount != chunkCount {
		msg := "分块信息错误: " + err.Error()
		log.Print(msg)
		res.Code = -1
		res.Msg = msg
		return
	}

	// TODO: 合并文件
	fs, _ := strconv.Atoi(fileSize)

	fileMeta := dbclient.FileMeta {
		FileSha1:   fileHash,
		FileName:   fileName,
		FileSize:   int64(fs),
		FileAddr:   "/data/"+upLoadID+"/"+chunkIndex,
	}

	dbclient.OnFileUploadFinished(fileMeta)
	dbclient.OnUserFileUploadFinished(username, fileMeta)
}

// 上传取消
func CancelUploadPartHandler(w http.ResponseWriter, r *http.Request) {
	// 删除存在的分块文件
	// 删除 redis 缓存状态
	// 更新 mysql 文件status
}

// 上传状态查询
func MultipartUploadStatusHandler(w http.ResponseWriter, r *http.Request) {
	// 检查分块上传状态
	// 获取分块初始化信息
	// 获取已上传的分块信息
}
