package main

import (
	"douyin-project/service/feed/config"
	"douyin-project/service/feed/route"
	"douyin-project/service/feed/svcctx"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	c, err := config.Parse("./config/feed.yaml")
	if err != nil {
		panic(err)
	}

	serviceCtx := svcctx.NewServiceContext(c)
	h := server.Default(server.WithHostPorts(c.ServiceAddress))
	route.RegisterHandler(h, c, serviceCtx)
	h.Spin()

}
