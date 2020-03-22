package orm

import (
	"log"

	mydb "admin-server/service/db/conn"
)

// 通过用户名及密码完成注册
func UserSignup(username string, passwd string) (res SqlResult) {
	// 准备sql语句 避免引号组装sql，防止注入
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`,`user_pwd`, `auth_type`) values (?,?,?)")
	if err != nil {
		log.Println("Failed to insert, err:" + err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd, "user")
	if err != nil {
		log.Println("Failed to insert, err:" + err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
		res.Suc = true
		return
	}

	res.Suc = false
	res.Msg = "无记录更新"
	return
}

// 用户登录
func UserSignin(username string, passwd string) (res SqlResult) {
	stmt, err := mydb.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}

	pRows := mydb.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == passwd {
		res.Suc = true
		res.Data = true
		return
	}
	res.Suc = false
	res.Msg = "用户名不存在或密码错误"
	return
}

// 更新token
func UpdateToken(username string, token string) (res SqlResult) {
	stmt, err := mydb.DBConn().Prepare(
		"replace into tbl_user_token (`user_name`,`user_token`) values (?,?)")
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	res.Suc = true
	return
}

// 查询用户信息
func GetUserInfo(username string) (res SqlResult) {
	user := TableUser{}

	stmt, err := mydb.DBConn().Prepare(
		"select user_name,signup_at, auth_type from tbl_user where user_name=? limit 1")
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt,&user.AuthType)
	if err != nil {
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	res.Suc = true
	res.Data = user
	return
}

// 查询用户Token
func GetUserToken(username, token string) (res SqlResult) {
	userToken := struct {
		Username string
		Token    string
	}{}

	stmt, err := mydb.DBConn().Prepare(
		"select user_name, user_token from tbl_user_token where user_name=? limit 1")
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&userToken.Username, &userToken.Token)
	if err != nil {
		res.Suc = false
		res.Msg = err.Error()
		return
	}

	if userToken.Token != token {
		res.Suc = false
		res.Msg = "Token 不一致"
		return
	}
	res.Suc = true
	return
}

// 查询用户是否存在
func UserExist(username string) (res SqlResult) {
	stmt, err := mydb.DBConn().Prepare(
		"select 1 from tbl_user where user_name=? limit 1")
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	res.Suc = true
	res.Data = map[string]bool{
		"exists": rows.Next(),
	}
	return
}
