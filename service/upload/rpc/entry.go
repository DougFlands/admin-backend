package rpc

import (
	"context"
	"admin-server/service/upload/config"
	uploadProto "admin-server/service/upload/proto"
	"fmt"
)

type Upload struct {

}


func (u *Upload) UploadEntry(ctx context.Context, req *uploadProto.ReqEntry, res *uploadProto.Resp) error {
	entry := config.UploadEntry
	//res.Code = 0
	fmt.Print(entry)

	return nil
}

