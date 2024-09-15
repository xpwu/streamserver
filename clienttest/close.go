package clienttest

import (
	"context"
	"fmt"
	"github.com/xpwu/go-log/log"
	"github.com/xpwu/go-tinyserver/api"
	pushapi "github.com/xpwu/streamserver/httpapi"
	"time"
)

type closeRequest struct {
}

func (s *suite) APIClose(ctx context.Context, _ *closeRequest) *api.EmptyResponse {
	ctx, logger := log.WithCtx(ctx)
	logger.PushPrefix(fmt.Sprintf("apiClose"))

	pushtoken := s.Req.Header.Get(pushTokenKey)

	go func() {
		time.Sleep(1 * time.Second)
		pushSuite := pushapi.NewPushSuiteForElseSuite(s.Req)
		pushSuite.APIClose(log.NewContext(context.Background(), logger), &pushapi.Request{PushToken: pushtoken})
	}()

	return &api.EmptyResponse{}
}
