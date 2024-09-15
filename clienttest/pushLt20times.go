package clienttest

import (
	"context"
	"fmt"
	"github.com/xpwu/go-log/log"
	"github.com/xpwu/go-tinyserver/api"
	pushapi "github.com/xpwu/streamserver/httpapi"
	"time"
)

type pushLt20TimesReq struct {
	Times int `json:"times"`
	// push: ${Prefix}-${no.}
	Prefix string `json:"prefix"`
}

func (s *suite) APIPushLt20Times(ctx context.Context, request *pushLt20TimesReq) *api.EmptyResponse {
	ctx, logger := log.WithCtx(ctx)
	logger.PushPrefix(fmt.Sprintf("APIPushLt20Times"))
	logger.Info(fmt.Sprintf("prefix: %s, times: %d", request.Prefix, request.Times))
	if request.Times > 20 {
		request.Times = 20
	}

	pushtoken := s.Req.Header.Get(pushTokenKey)

	go func() {
		ctx := log.NewContext(context.Background(), logger)
		pushSuite := pushapi.NewPushSuiteForElseSuite(s.Req)
		for i := 0; i < request.Times; i++ {
			pushSuite.APIPush(ctx, pushapi.NewPushDateRequest(pushtoken,
				fmt.Sprintf("%s-%d", request.Prefix, i)))
			time.Sleep(1 * time.Second)
		}
	}()

	return &api.EmptyResponse{}
}
