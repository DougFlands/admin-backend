package main

import (
	"admin-server/config"
	"admin-server/service/apigateway/route"
)

func main()  {

	r := route.Router()
	r.Run(config.ApiGatewayServiceHost)
}
