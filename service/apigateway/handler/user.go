package handler

import (
	"context"
	userProto "admin-server/service/account/proto"
	"admin-server/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 用户注册请求
func DoSignUpHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	resp, err := userClient.Signup(context.TODO(), &userProto.ReqSignup{
		UserName: username,
		PassWord: passwd,
	})

	if err != nil {
		log.Print(err.Error())
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, util.NewRespMsg(int(resp.Code), resp.Msg, nil))
}

// 用户登录请求
func DoSignInHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	resp, err := userClient.Signin(context.TODO(), &userProto.ReqSignin{
		UserName: username,
		PassWord: passwd,
	})

	if err != nil {
		log.Print(err.Error())
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, util.NewRespMsg(int(resp.Code), resp.Msg, resp.Data))
}

// 查询用户信息
func UserInfoHandler(c *gin.Context) {
	username, _ := c.Request.Cookie("username")
	token, _ := c.Request.Cookie("token")

	resp, err := userClient.UserInfo(context.TODO(), &userProto.ReqUserInfo{
		UserName: username.Value,
		Token:    token.Value,
	})

	if err != nil {
		log.Print(err.Error())
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, util.NewRespMsg(int(resp.Code), resp.Msg, resp.Data))
}
