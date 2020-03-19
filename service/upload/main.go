package main

import (
	"admin-server/config"
	dbproxy "admin-server/service/db/client"
	"admin-server/service/upload/route"
	"github.com/micro/go-micro"
	"log"
)

func startRpcService() {
	service := micro.NewService(
		micro.Registry(config.ConsulReg),
		micro.Name("go.micro.service.upload"),
	)

	service.Init()

	// 初始化dbproxy client
	dbproxy.Init(service)

	//uploadProto.RegisterUploadServiceHandler(service.Server(), new(rpc.Upload))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func startApiService() {
	router := route.Router()
	router.Run(config.UploadServiceHost)
}

func main() {
	go startRpcService()
	startApiService()
}
