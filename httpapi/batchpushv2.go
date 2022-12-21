package httpapi

import (
  "context"
  "fmt"
  "github.com/xpwu/go-log/log"
  "github.com/xpwu/go-stream/push/core"
)

type batchPushV2Request struct {
  Requests []pushdateRequest `json:"reqs"`
}

func (s *suite) APIBatchPushV2(ctx context.Context, request *batchPushV2Request) *batchResponse {
  ctx, logger := log.WithCtx(ctx)
  logger.PushPrefix(fmt.Sprintf("api push count %d", len(request.Requests)))

  res := &batchResponse{States: make([]int, 0, len(request.Requests))}

  for _,req := range request.Requests {
    clientConn, st := core.GetClientConn(ctx, req.PushToken)
    if st != core.Success {
      res.States = append(res.States, st.ToInt())
      continue
    }

    st = core.PushDataToClient(ctx, clientConn, []byte(req.Data))
    res.States = append(res.States, st.ToInt())
  }

  return res
}
