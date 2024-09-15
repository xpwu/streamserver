package httpapi

import "github.com/xpwu/go-tinyserver/api"

type Suite = suite

func NewPushSuiteForElseSuite(req *api.Request) *Suite {
	return &suite{
		PostJsonSetUpper: api.PostJsonSetUpper{Req: req},
	}
}

type PushDateRequest = pushdateRequest

func NewPushDateRequest(pushToken string, data string) *PushDateRequest {
	return &pushdateRequest{
		request{PushToken: pushToken},
		data,
	}
}

type Response = response

type Request = request
