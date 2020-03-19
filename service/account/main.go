package main

import (
	"admin-server/config"
	"admin-server/service/account/handler"
	proto "admin-server/service/account/proto"
	dbproxy "admin-server/service/db/client"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Registry(config.ConsulReg),
		micro.Name("go.micro.service.user"),
	)

	service.Init()
	// 初始化dbproxy client
	dbproxy.Init(service)

	proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
