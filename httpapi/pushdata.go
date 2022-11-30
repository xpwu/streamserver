package httpapi

import (
  "context"
  "github.com/xpwu/go-log/log"
  "github.com/xpwu/go-stream/push/core"
)

type pushdateRequest struct {
  request
  Data string `json:"data"`
}

func (s *suite) APIPush(ctx context.Context, request *pushdateRequest) *response {
  ctx, logger := log.WithCtx(ctx)
  logger.PushPrefix("api push: " + request.PushToken)

  clientConn, st := core.GetClientConn(ctx, request.PushToken)
  if st != core.Success {
    return &response{State: st.ToInt()}
  }

  st = core.PushDataToClient(ctx, clientConn, []byte(request.Data))

  return &response{State: st.ToInt()}
}
