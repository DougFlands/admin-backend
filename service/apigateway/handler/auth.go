package handler

import (
	dbclient "admin-server/service/db/client"
	"admin-server/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IsTokenValid(username string, token string) bool {
	_, err := dbclient.GetUserToken(username, token)
	if err == "" {
		return true
	}
	log.Print(err)
	return false
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _ := c.Request.Cookie("username")
		token, _ := c.Request.Cookie("token")

		//验证登录token是否有效
		if len(username.Value) < 3 || !IsTokenValid(username.Value, token.Value) {
			// w.WriteHeader(http.StatusForbidden)
			// token校验失败则跳转到登录页面
			c.Abort()
			resp := util.NewRespMsg(
				-1,
				"token无效",
				nil,
			)
			c.JSON(http.StatusOK, resp)
			return
		}
		c.Next()
	}
}
