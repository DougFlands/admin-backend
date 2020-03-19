package orm

import (
	"log"
	"time"

	mydb "admin-server/service/db/conn"
)

// 更新用户文件表
func OnUserFileUploadFinished(username, filehash, filename string, filesize int64) (res SqlResult) {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user_file (`user_name`,`file_sha1`,`file_name`,`file_size`,`upload_at`) values (?,?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, filehash, filename, filesize, time.Now())
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	res.Suc = true
	return
}

// 查询用户文件列表
func QueryUserFileMetas(username string, limit int64) (res SqlResult) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1,file_name,file_rename,file_size,upload_at,last_update from tbl_user_file where user_name=? and status=0 limit ?")
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, limit)
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}

	var userFiles []TableUserFile
	for rows.Next() {
		ufile := TableUserFile{}
		err = rows.Scan(&ufile.FileHash, &ufile.FileName, &ufile.FileRename, &ufile.FileSize,
			&ufile.UploadAt, &ufile.LastUpdated)
		if err != nil {
			log.Println(err.Error())
			break
		}
		userFiles = append(userFiles, ufile)
	}
	res.Suc = true
	res.Data = userFiles
	return
}

// 标记文件删除
func DeleteUserFile(username, filehash string) (res SqlResult) {
	stmt, err := mydb.DBConn().Prepare(
		"update tbl_user_file set status=2 where user_name=? and file_sha1=? limit 1")
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, filehash)
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	res.Suc = true
	return
}

// 重命名文件
func RenameFileName(username, filehash, rename string) (res SqlResult) {
	stmt, err := mydb.DBConn().Prepare(
		"update tbl_user_file set file_rename=? where user_name=? and file_sha1=? limit 1")
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(rename, username, filehash)
	if err != nil {
		log.Println(err.Error())
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	res.Suc = true
	return
}

// 用户单个文件信息
func QueryUserFileMeta(username string, filehash string) (res SqlResult) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1,file_name,file_rename,file_size,upload_at,last_update from tbl_user_file where user_name=? and file_sha1=?  limit 1")
	if err != nil {
		res.Suc = false
		res.Msg = err.Error()
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, filehash)
	if err != nil {
		res.Suc = false
		res.Msg = err.Error()
		return
	}

	ufile := TableUserFile{}
	if rows.Next() {
		err = rows.Scan(&ufile.FileHash, &ufile.FileName, &ufile.FileRename, &ufile.FileSize,
			&ufile.UploadAt, &ufile.LastUpdated)
		if err != nil {
			log.Println(err.Error())
			res.Suc = false
			res.Msg = err.Error()
			return
		}
	}

	res.Suc = true
	res.Data = ufile
	return
}
