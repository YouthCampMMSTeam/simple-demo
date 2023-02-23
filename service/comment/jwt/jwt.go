package jwt

import (
	"context"
	"douyin-project/common/token"
	"douyin-project/service/comment/config"
	"douyin-project/service/comment/svcctx"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

// func InitJwt() {
func NewJwtjwtMiddleware(c *config.Config, svc *svcctx.ServiceContext) *jwt.HertzJWTMiddleware {
	IdentityKey := c.IdentityKey
	jwtKey := c.JwtKey

	jwtMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "douyin-project",
		Key:           []byte(jwtKey),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"message": "success",
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			//因为这里没有登陆接口所以就没有实现。。
			return struct{}{}, nil
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			// claims := jwt.ExtractClaims(ctx, c)
			// return &model.User{
			// 	UserName: claims[IdentityKey].(string),
			// }

			//我猜测上下是一致的 即IdentityHandler和PayloadFunc
			claims := jwt.ExtractClaims(ctx, c)
			return &token.TokenUser{
				Id: int64(claims[IdentityKey].(float64)),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// if v, ok := data.(*model.User); ok {
			// 	return jwt.MapClaims{
			// 		IdentityKey: v.UserName,
			// 	}
			// }
			// return jwt.MapClaims{}

			if v, ok := data.(*token.TokenUser); ok {
				return jwt.MapClaims{
					IdentityKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt HTTPStatusMessageFunc err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
	return jwtMiddleware
}
