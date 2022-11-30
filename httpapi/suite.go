package httpapi

import (
  "github.com/xpwu/go-tinyserver/api"
)

type request struct {
  PushToken string `json:"pushtoken"`
}

type response struct {
  State int `json:"st"`
}

type suite struct {
  api.PostJsonSetUpper
  api.PostJsonTearDowner
  api.RootURIMapper
}

func AddAPI() {
  api.Add(func() api.Suite {
    return &suite{}
  })
}
