package main

import (
	"admin-server/config"
	"log"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	dbConn "admin-server/service/db/conn"
	dbProxy "admin-server/service/db/proto"
	dbRpc "admin-server/service/db/rpc"
)

func startRpcService() {
	service := micro.NewService(
		micro.Name("go.micro.service.dbproxy"), // 在注册中心中的服务名称
		micro.Registry(config.ConsulReg),
		micro.RegisterTTL(time.Second*10),      // 声明超时时间, 避免consul不主动删掉已失去心跳的服务节点
		micro.RegisterInterval(time.Second*5),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {}),
	)

	// 初始化db connection
	dbConn.InitDBConn()

	dbProxy.RegisterDBProxyServiceHandler(service.Server(), new(dbRpc.DBProxy))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func main() {
	startRpcService()
}