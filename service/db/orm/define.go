package orm

import "database/sql"

// 文件表
type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

// 用户表
type TableUser struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
	AuthType 	 string
}

// 用户文件表
type TableUserFile struct {
	UserName    string
	FileHash    string
	FileName    string
	FileRename    string
	FileSize    int64
	UploadAt    string
	LastUpdated string
}

// sql函数执行的结果
type SqlResult struct {
	Suc  bool
	Code int
	Msg  string
	Data interface{}
}
