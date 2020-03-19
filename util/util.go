package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"hash"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Sha1Stream struct {
	_sha1 hash.Hash
}

func (obj *Sha1Stream) Update(data []byte) {
	if obj._sha1 == nil {
		obj._sha1 = sha1.New()
	}
	obj._sha1.Write(data)
}

func (obj *Sha1Stream) Sum() string {
	return hex.EncodeToString(obj._sha1.Sum([]byte("")))
}

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func FileSha1(file *os.File) string {
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum(nil))
}

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func FileMD5(file *os.File) string {
	_md5 := md5.New()
	io.Copy(_md5, file)
	return hex.EncodeToString(_md5.Sum(nil))
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}


// 生成token
func GenToken(username string) string {
	nowData := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := MD5([]byte(username + nowData + "_token"))
	return tokenPrefix + nowData[:8]
}


type RpcResp struct {
	Code int
	Msg string
	Data []byte
}

// 构建rpc通讯参数
func InitRpc() RpcResp  {
	var data []byte
	rpc := RpcResp {
		Code: 0,
		Msg: "",
		Data: data,
	}
	return rpc
}

// 设置 resp
func DeferRpc(res *RpcResp, c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

	if res.Code == 0 {
		c.JSON(http.StatusOK, NewRespMsg(0, res.Msg, res.Data))
	} else {
		c.JSON(http.StatusOK, NewRespMsg(-1,res.Msg, res.Data))
	}
}


