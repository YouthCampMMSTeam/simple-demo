package main

import (
	"douyin-project/service/user/config"
	"douyin-project/service/user/route"
	"douyin-project/service/user/svcctx"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	c, err := config.Parse("./config/user.yaml")
	if err != nil {
		panic(err)
	}

	svc := svcctx.NewServiceContext(c)

	h := server.Default(server.WithHostPorts(c.ServiceAddress))
	route.RegisterHandler(h, c, svc)
	h.Spin()
}
