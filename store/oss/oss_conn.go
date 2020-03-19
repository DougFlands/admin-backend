package oss

import (
	"context"
	"admin-server/config"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"time"
)

var ossCli *cos.Client

// 创建oss client
func Client() *cos.Client{
	if ossCli != nil {
		return ossCli
	}

	u, _ := url.Parse(config.OSSPath)
	b := &cos.BaseURL{BucketURL: u}

	 ossCli = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.OSSAccessKeyID,
			SecretKey: config.OSSAccessKeySecret,
		},
	})

	return ossCli
}

// 获取预签名URL
func DownloadPresignedUrl(key string) string {
	presignedURL, err := Client().Object.GetPresignedURL(
		context.Background(),
		http.MethodGet,
		key,
		config.OSSAccessKeyID,
		config.OSSAccessKeySecret,
		time.Hour,
		nil)
	if err != nil {
		panic(err)
	}
	return presignedURL.String()
}
