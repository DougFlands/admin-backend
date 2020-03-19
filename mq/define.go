package mq

type TransferData struct {
	FileName string
	FileHash string // 文件hash
	CurLocation string // 存在临时目录地址
	DestLocation string // 存在OSS哪里
}

