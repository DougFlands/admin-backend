package client

import (
	"context"
	"encoding/json"
	"github.com/mitchellh/mapstructure"

	"github.com/micro/go-micro"

	"admin-server/service/db/orm"
	dbProto "admin-server/service/db/proto"
)

// FileMeta : 文件元信息结构
type FileMeta struct {
	FileSha1   string
	FileName   string
	FileReName string
	FileSize   int64
	FileAddr   string
	UploadAt   string
}

var (
	dbClient dbProto.DBProxyService
)

func Init(service micro.Service) {
	// 初始化一个dbproxy服务的客户端
	dbClient = dbProto.NewDBProxyService("go.micro.service.dbproxy", service.Client())
}


func TableFileToFileMeta(tfile orm.TableFile) FileMeta {
	return FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		FileAddr: tfile.FileAddr.String,
	}
}

// 向dbproxy请求执行action
func execAction(funcName string, paramJson []byte) (*dbProto.RespExec, error) {
	return dbClient.ExecuteAction(context.TODO(), &dbProto.ReqExec{
		Action: []*dbProto.SingleAction{
			{
				Name:   funcName,
				Params: paramJson,
			},
		},
	})
}



// 提取DB返回的错误信息
func formatDbResp(dbResp *dbProto.RespExec, err error) (*orm.SqlResult, string) {
	// 转换rpc返回的结果
	parseBody := func (resp *dbProto.RespExec) *orm.SqlResult {
		if resp == nil || resp.Data == nil {
		return nil
	}
		var resList []orm.SqlResult
		_ = json.Unmarshal(resp.Data, &resList)
		if len(resList) > 0 {
		return &resList[0]
	}
		return nil
	}
	result := parseBody(dbResp)
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	if !result.Suc {
		errMsg = dbResp.Msg
	}
	return result, errMsg
}

func ToTableUser(src interface{}) orm.TableUser {
	user := orm.TableUser{}
	mapstructure.Decode(src, &user)
	return user
}

func ToTableFile(src interface{}) orm.TableFile {
	file := orm.TableFile{}
	mapstructure.Decode(src, &file)
	return file
}

func ToTableFiles(src interface{}) []orm.TableFile {
	file := []orm.TableFile{}
	mapstructure.Decode(src, &file)
	return file
}

func ToTableUserFile(src interface{}) orm.TableUserFile {
	ufile := orm.TableUserFile{}
	mapstructure.Decode(src, &ufile)
	return ufile
}

func ToTableUserFiles(src interface{}) []orm.TableUserFile {
	ufile := []orm.TableUserFile{}
	mapstructure.Decode(src, &ufile)
	return ufile
}

func GetFileMeta(filehash string) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{filehash})
	res, err := execAction("/file/GetFileMeta", uInfo)
	return formatDbResp(res, err)
}

func GetFileMetaList(limitCnt int) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{limitCnt})
	res, err := execAction("/file/GetFileMetaList", uInfo)
	return formatDbResp(res, err)
}

func OnFileUploadFinished(fmeta FileMeta) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.FileAddr})
	res, err := execAction("/file/OnFileUploadFinished", uInfo)
	return formatDbResp(res, err)
}

func UpdateFileLocation(filehash, location string) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{filehash, location})
	res, err := execAction("/file/UpdateFileLocation", uInfo)
	return formatDbResp(res, err)
}

func UserSignup(username, encPasswd string) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{username, encPasswd})
	res, err := execAction("/user/UserSignup", uInfo)
	return formatDbResp(res, err)
}

func UserSignin(username, encPasswd string) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{username, encPasswd})
	res, err := execAction("/user/UserSignin", uInfo)
	return formatDbResp(res, err)
}

func GetUserInfo(username string) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{username})
	res, err := execAction("/user/GetUserInfo", uInfo)
	return formatDbResp(res, err)
}

func GetUserToken(username , token string) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{username, token})
	res, err := execAction("/user/GetUserToken", uInfo)
	return formatDbResp(res, err)
}


func UserExist(username string) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{username})
	res, err := execAction("/user/UserExist", uInfo)
	return formatDbResp(res, err)
}

func UpdateToken(username, token string) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{username, token})
	res, err := execAction("/user/UpdateToken", uInfo)
	return formatDbResp(res, err)
}

func QueryUserFileMeta(username, filehash string) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{username, filehash})
	res, err := execAction("/ufile/QueryUserFileMeta", uInfo)
	return formatDbResp(res, err)
}

func QueryUserFileMetas(username string, limit int) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{username, limit})
	res, err := execAction("/ufile/QueryUserFileMetas", uInfo)
	return formatDbResp(res, err)
}

// 新增/更新文件元信息到mysql中
func OnUserFileUploadFinished(username string, fmeta FileMeta) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{username, fmeta.FileSha1,
		fmeta.FileName, fmeta.FileSize})
	res, err := execAction("/ufile/OnUserFileUploadFinished", uInfo)
	return formatDbResp(res, err)
}

func RenameFileName(username, filehash, filename string) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{username, filehash, filename})
	res, err := execAction("/ufile/RenameFileName", uInfo)
	return formatDbResp(res, err)
}

func DeleteUserFile(username, filehash string) (*orm.SqlResult, string) {
	uInfo, _ := json.Marshal([]interface{}{username, filehash})
	res, err := execAction("/ufile/DeleteUserFile", uInfo)
	return formatDbResp(res, err)
}
