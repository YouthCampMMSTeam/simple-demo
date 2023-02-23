package route

import (
	"douyin-project/service/comment/config"
	"douyin-project/service/comment/jwt"
	"douyin-project/service/comment/route/handler"
	"douyin-project/service/comment/svcctx"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterHandler(h *server.Hertz, c *config.Config, serverCtx *svcctx.ServiceContext) {
	jwtMiddleware := jwt.NewJwtjwtMiddleware(c, serverCtx)
	douyinGroup := h.Group("/douyin", jwtMiddleware.MiddlewareFunc())

	douyinGroup.POST("/comment/action/", handler.CommentActionHandler(serverCtx))
	douyinGroup.GET("/comment/list/", handler.CommentListHandler(serverCtx))
}
