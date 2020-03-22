package handler

import (
	"admin-server/config"
	userProto "admin-server/service/account/proto"
	dbclient "admin-server/service/db/client"
	"admin-server/util"
	"context"
	"encoding/json"
)

type User struct {
}

// 初始化 Resp
func initResp(res *userProto.Resp) {
	res.Data = []byte{}
	res.Msg = ""
	res.Code = 0
}

// 用户注册请求
func (u *User) Signup(ctx context.Context, req *userProto.ReqSignup, res *userProto.Resp) error {
	username := req.UserName
	passwd := req.PassWord

	initResp(res)

	if len(username) < 3 || len(passwd) < 3 {
		res.Code = -1
		res.Msg = "用户名及密码需大于三位"
		return nil
	}

	encPasswd := util.Sha1([]byte(passwd + config.PwdSalt))
	_, err := dbclient.UserSignup(username, encPasswd)
	if err == "" {
		res.Code = 0
		res.Msg = "注册成功"
	} else {
		res.Code = -1
		res.Msg = "注册失败: " + err
	}

	return nil
}

// 用户登录请求
func (u *User) Signin(ctx context.Context, req *userProto.ReqSignin, res *userProto.Resp) error {
	username := req.UserName
	passwd := req.PassWord
	initResp(res)

	if username == "" {
		res.Code = -1
		res.Msg = "用户名不能为空"
		return nil
	}

	encPasswd := util.Sha1([]byte(passwd + config.PwdSalt))
	_, err := dbclient.UserSignin(username, encPasswd)
	if err != "" {
		res.Code = -1
		res.Msg = err
		return nil
	}

	token := util.GenToken(username)
	_, err = dbclient.UpdateToken(username, token)
	if err != "" {
		res.Code = -1
	}

	res.Code = 0

	data, _ := json.Marshal(map[string]string{
		"Username": username,
		"Token":    token,
	})

	res.Data = data

	return nil
}

// 查询用户信息
func (u *User) UserInfo(ctx context.Context, req *userProto.ReqUserInfo, res *userProto.Resp) error {
	username := req.UserName
	initResp(res)

	dbResp, err := dbclient.GetUserInfo(username)
	if err != "" {
		res.Code = -1
		res.Msg = "数据库获取用户信息失败"
		return nil
	}

	data, _ := json.Marshal(dbResp.Data)

	res.Data = data
	return nil
}
