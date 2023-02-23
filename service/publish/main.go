package main

import (
	"douyin-project/service/publish/config"
	"douyin-project/service/publish/route"
	"douyin-project/service/publish/svcctx"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	c, err := config.Parse("./config/publish.yaml")
	if err != nil {
		panic(err)
	}

	serviceCtx := svcctx.NewServiceContext(c)
	h := server.Default(server.WithHostPorts(c.ServiceAddress))
	route.RegisterHandler(h, c, serviceCtx)
	h.Spin()
}
