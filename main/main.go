package main

import (
	"github.com/xpwu/go-cmd/arg"
	"github.com/xpwu/go-cmd/cmd"
	_ "github.com/xpwu/go-cmd/cmd/printconf"
	"github.com/xpwu/go-stream/lencontent"
	"github.com/xpwu/go-stream/push"
	"github.com/xpwu/go-stream/websocket"
	"github.com/xpwu/go-tinyserver/http"
	"github.com/xpwu/streamserver/clienttest"
	"github.com/xpwu/streamserver/httpapi"
)

func main() {

	cmd.RegisterCmd(cmd.DefaultCmdName, "start server", func(args *arg.Arg) {

		arg.ReadConfig(args)
		args.Parse()

		websocket.Start()
		lencontent.Start()
		push.Start()

		httpapi.AddAPI()
		clienttest.AddAPI()
		http.Start()

		// block
		block := make(chan struct{})
		<-block
	})

	cmd.Run()
}
