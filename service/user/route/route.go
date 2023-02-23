package route

import (
	"douyin-project/service/user/config"
	"douyin-project/service/user/jwt"
	"douyin-project/service/user/route/handler"
	"douyin-project/service/user/svcctx"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterHandler(h *server.Hertz, c *config.Config, serverCtx *svcctx.ServiceContext) {

	jwtMiddleware := jwt.NewJwtjwtMiddleware(c, serverCtx)

	douyinGroup := h.Group("/douyin")
	douyinGroup.POST("/user/login/", jwtMiddleware.LoginHandler)
	douyinGroup.GET("/user/register/", handler.UserRegisterHandler(serverCtx), jwtMiddleware.LoginHandler) //可以多个handler 后面增加login来处理token

	//中间件应该也是handler
	douyinGroup.GET("/user/", jwtMiddleware.MiddlewareFunc(), handler.UserInfoHandler(serverCtx))

}
