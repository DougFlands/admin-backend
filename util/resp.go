package util

import (
	"encoding/json"
	"log"
)

// RespMsg : http响应数据的通用结构
type RespMsg struct {
	Code int
	Msg  string
	Data interface{}
}

// NewRespMsg : 生成response对象
func NewRespMsg(code int, msg string, data []byte) *RespMsg {
	result := make(map[string]interface{})
	if data != nil {
		result = ConversionProtobufData(data)
	}
	return &RespMsg{
		Code: code,
		Msg:  msg,
		Data: result,
	}
}

// JSONBytes : 对象转json格式的二进制数组
func (resp *RespMsg) JSONBytes() []byte {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return r
}

// JSONString : 对象转json格式的string
func (resp *RespMsg) JSONString() string {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return string(r)
}

// 反 json 化 protobuf 给的 data 数据
func ConversionProtobufData(data []byte) map[string]interface{} {
	var mapResult map[string]interface{}
	if err := json.Unmarshal(data, &mapResult); err != nil {
		log.Print(err)
	}
	return mapResult
}
