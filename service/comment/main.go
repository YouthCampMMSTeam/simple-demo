package main

import (
	"douyin-project/service/comment/config"
	"douyin-project/service/comment/route"
	"douyin-project/service/comment/svcctx"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	c, err := config.Parse("./config/comment.yaml")
	if err != nil {
		panic(err)
	}

	serviceCtx := svcctx.NewServiceContext(c)
	h := server.Default(server.WithHostPorts(c.ServiceAddress))
	route.RegisterHandler(h, c, serviceCtx)
	h.Spin()

}
