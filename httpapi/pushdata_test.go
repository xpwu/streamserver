package httpapi

import (
  "encoding/json"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestFlatten(t *testing.T) {
  r := &pushdateRequest{
    request: request{PushToken: "abc"},
    Data:    "this is data",
  }

  d,_ := json.Marshal(r)

  assert.Equal(t, "{\"tk\":\"abc\",\"data\":\"this is data\"}", string(d))
}
