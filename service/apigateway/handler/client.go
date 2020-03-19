package handler


import (
	"admin-server/config"
	userProto "admin-server/service/account/proto"
	dbproxy "admin-server/service/db/client"
	uploadProto "admin-server/service/upload/proto"
	"github.com/micro/go-micro"
)

var (
	userClient userProto.UserService
	uploadClient uploadProto.UploadService
)

func init() {
	service := micro.NewService(micro.Registry(config.ConsulReg))
	service.Init()

	dbproxy.Init(service)

	// 初始化
	userClient = userProto.NewUserService("go.micro.service.user", service.Client())
	uploadClient = uploadProto.NewUploadService("go.micro.service.upload", service.Client())
}
