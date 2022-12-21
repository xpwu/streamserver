package httpapi

import (
  "context"
  "fmt"
  "github.com/xpwu/go-log/log"
  "github.com/xpwu/go-stream/push/core"
)

type batchCloseRequest struct {
  PushTokens []string `json:"pushtokens"`
}

func (s *suite) APIBatchClose(ctx context.Context, request *batchCloseRequest) *batchResponse {
  ctx, logger := log.WithCtx(ctx)
  logger.PushPrefix(fmt.Sprintf("api push count %d", len(request.PushTokens)))

  res := &batchResponse{States: make([]int, 0, len(request.PushTokens))}

  for _,token := range request.PushTokens {
    clientConn, st := core.GetClientConn(ctx, token)
    if st == core.Success {
      core.CloseClientConn(ctx, clientConn)
    }

    res.States = append(res.States, st.ToInt())
  }

  return res
}
