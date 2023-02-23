package route

import (
	"douyin-project/service/publish/config"
	"douyin-project/service/publish/jwt"
	"douyin-project/service/publish/route/handler"
	"douyin-project/service/publish/svcctx"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterHandler(h *server.Hertz, c *config.Config, serverCtx *svcctx.ServiceContext) {
	jwtMiddleware := jwt.NewJwtjwtMiddleware(c, serverCtx)
	douyinGroup := h.Group("/douyin", jwtMiddleware.MiddlewareFunc())

	douyinGroup.POST("/publish/action/", handler.PublishActionHandler(serverCtx))
	douyinGroup.GET("/publish/list/", handler.PublishListHandler(serverCtx))
}
