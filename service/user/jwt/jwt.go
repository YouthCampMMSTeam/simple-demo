package jwt

import (
	"context"
	"douyin-project/common/token"
	"douyin-project/service/user/config"
	"douyin-project/service/user/logic"
	"douyin-project/service/user/svcctx"
	"douyin-project/service/user/types"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

// var (
// JwtMiddleware *jwt.HertzJWTMiddleware
// 	IdentityKey   = "identity"
// 	// jwtKey        = "secret key"
// 	jwtKey = "ae0536f9-6450-4606-8e13-5a19ed505da0"
// )

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
			//TODO 注意字段限制
			// var loginStruct struct {
			// 	name     string `form:"name" json:"name" query:"name" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			// 	Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			// }
			// if err := c.BindAndValidate(&loginStruct); err != nil {
			// 	return nil, err
			// }
			req := &types.UserLoginReq{}
			if err := c.BindAndValidate(req); err != nil {
				return nil, err
			}
			fmt.Printf("req %+v\n", req)

			l := logic.NewUserLogic(svc)
			resp, err := l.UserLogin(ctx, &types.UserLoginLogicReq{
				Name:     req.Name,
				Password: req.Password,
			})
			if err != nil {
				return nil, err
			}
			return &token.TokenUser{
				Id: resp.UserId,
			}, nil
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
