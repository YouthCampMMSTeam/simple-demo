package main

import (
	"douyin-project/service/favorite/config"
	"douyin-project/service/favorite/route"
	"douyin-project/service/favorite/svcctx"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	c, err := config.Parse("./config/favorite.yaml")
	if err != nil {
		panic(err)
	}

	serviceCtx := svcctx.NewServiceContext(c)
	h := server.Default(server.WithHostPorts(c.ServiceAddress))
	route.RegisterHandler(h, c, serviceCtx)
	h.Spin()

}
