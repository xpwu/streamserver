package clienttest

import "github.com/xpwu/go-tinyserver/api"

const pushTokenKey = "pushtoken"

type suite struct {
	api.PostJsonSetUpper
	api.PostJsonTearDowner
}

func (s *suite) MappingPreUri() string {
	return "/clienttest"
}

func AddAPI() {
	api.Add(func() api.Suite {
		return &suite{}
	})
}
