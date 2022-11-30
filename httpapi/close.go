package httpapi

import (
  "context"
  "github.com/xpwu/go-log/log"
  "github.com/xpwu/go-stream/push/core"
)

func (s *suite) APIClose(ctx context.Context, request *request) *response {
  ctx, logger := log.WithCtx(ctx)
  logger.PushPrefix("api close: " + request.PushToken)

  clientConn, st := core.GetClientConn(ctx, request.PushToken)
  if st != core.Success {
    return &response{State: st.ToInt()}
  }

  core.CloseClientConn(ctx, clientConn)

  return &response{State: core.Success.ToInt()}
}
