package route

import (
	"douyin-project/service/favorite/config"
	"douyin-project/service/favorite/jwt"
	"douyin-project/service/favorite/route/handler"
	"douyin-project/service/favorite/svcctx"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterHandler(h *server.Hertz, c *config.Config, serverCtx *svcctx.ServiceContext) {
	jwtMiddleware := jwt.NewJwtjwtMiddleware(c, serverCtx)
	douyinGroup := h.Group("/douyin", jwtMiddleware.MiddlewareFunc())

	douyinGroup.POST("/favorite/action/", handler.FavoriteActionHandler(serverCtx))
	douyinGroup.GET("/favorite/list/", handler.FavoriteListHandler(serverCtx))
}
