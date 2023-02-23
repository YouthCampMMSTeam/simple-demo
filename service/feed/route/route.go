package route

import (
	"douyin-project/service/feed/config"
	"douyin-project/service/feed/jwt"
	"douyin-project/service/feed/route/handler"
	"douyin-project/service/feed/svcctx"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterHandler(h *server.Hertz, c *config.Config, serverCtx *svcctx.ServiceContext) {
	jwtMiddleware := jwt.NewJwtjwtMiddleware(c, serverCtx)
	douyinGroup := h.Group("/douyin", jwtMiddleware.MiddlewareFunc())
	douyinGroup.GET("/feed", handler.FeedHandler(serverCtx))
}
