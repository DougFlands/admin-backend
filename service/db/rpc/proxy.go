package rpc

import (
	"bytes"
	"context"
	"encoding/json"

	"admin-server/service/db/mapper"
	"admin-server/service/db/orm"
	dbProxy "admin-server/service/db/proto"
)

// DBProxy : DBProxy结构体
type DBProxy struct{}

// ExecuteAction : 请求执行sql函数
func (db *DBProxy) ExecuteAction(ctx context.Context, req *dbProxy.ReqExec, res *dbProxy.RespExec) error {
	resList := make([]orm.SqlResult, len(req.Action))

	for idx, singleAction := range req.Action {
		var params []interface{}
		dec := json.NewDecoder(bytes.NewReader(singleAction.Params))
		dec.UseNumber()
		// 避免int/int32/int64等自动转换为float64
		if err := dec.Decode(&params); err != nil {
			resList[idx] = orm.SqlResult{
				Suc: false,
				Msg: "请求参数有误",
			}
			continue
		}

		for k, v := range params {
			if _, ok := v.(json.Number); ok {
				params[k], _ = v.(json.Number).Int64()
			}
		}

		// 默认串行执行sql函数
		execRes, err := mapper.FuncCall(singleAction.Name, params...)
		if err != nil {
			resList[idx] = orm.SqlResult{
				Suc: false,
				Msg: "函数调用有误",
			}
			continue
		}
		resList[idx] = execRes[0].Interface().(orm.SqlResult)
	}

	res.Data, _ = json.Marshal(resList)
	return nil
}
