package httpapi

import (
  "context"
  "fmt"
  "github.com/xpwu/go-log/log"
  "github.com/xpwu/go-stream/push/core"
)

type batchPushV1Request struct {
  PushTokens []string `json:"pushtokens"`
  Data string `json:"data"`
}

func (s *suite) APIBatchPushV1(ctx context.Context, request *batchPushV1Request) *batchResponse {
  ctx, logger := log.WithCtx(ctx)
  logger.PushPrefix(fmt.Sprintf("api push count %d", len(request.PushTokens)))

  res := &batchResponse{States: make([]int, 0, len(request.PushTokens))}

  for _,token := range request.PushTokens {
    clientConn, st := core.GetClientConn(ctx, token)
    if st != core.Success {
      res.States = append(res.States, st.ToInt())
      continue
    }

    st = core.PushDataToClient(ctx, clientConn, []byte(request.Data))
    res.States = append(res.States, st.ToInt())
  }

  return res
}
