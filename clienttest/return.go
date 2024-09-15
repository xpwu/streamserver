package clienttest

import (
	"context"
	"fmt"
	"github.com/xpwu/go-log/log"
)

type returnReq struct {
	Data string `json:"data"`
}

type returnRes struct {
	Ret string `json:"ret"`
}

func (s *suite) APIReturn(ctx context.Context, request *returnReq) *returnRes {
	ctx, logger := log.WithCtx(ctx)
	logger.PushPrefix(fmt.Sprintf("apiReturn"))
	logger.Info(request.Data)

	return &returnRes{Ret: request.Data}
}
