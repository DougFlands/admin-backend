package main

import (
	"admin-server/config"
	"admin-server/mq"
	dbclient "admin-server/service/db/client"
	"admin-server/store/oss"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ProcessTransfer(msg []byte) {
	log.Print("开始上传至oss")
	// msg 转成对应的结构
	pubData := mq.TransferData{}
	err := json.Unmarshal(msg, &pubData)

	if err != nil {
		log.Print(err.Error())
		return
	}

	// 读取文件
	filed, err := os.Open(pubData.CurLocation)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	name := pubData.FileName + "-" + pubData.FileHash

	// 上传到oss
	_, err = oss.Client().Object.Put(context.Background(), name, filed, nil)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	//更新数据库
	_, error := dbclient.UpdateFileLocation(
		pubData.FileHash,
		pubData.DestLocation)

	if error == "" {
		log.Print("上传成功")
	} else {
		log.Print("上传失败")
	}
}

func main() {
	log.Println("监听转移任务队列")
	mq.StartConsume(config.KafkaTopic, ProcessTransfer)
}
