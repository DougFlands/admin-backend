package route

import (
	"admin-server/service/apigateway/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	//router.Static("/static", "./static")
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Range", "x-requested-with", "content-Type"},
		ExposeHeaders: []string{"Content-Length", "Accept-Ranges", "Content-Range", "Content-Disposition"},
		AllowCredentials: true,
	}))

	router.POST("/api/user/signin", handler.DoSignInHandler)

	router.Use(handler.Authorize())

	router.POST("/api/user/signup", handler.DoSignUpHandler)
	router.POST("/api/user/info", handler.UserInfoHandler)

	return router
}

