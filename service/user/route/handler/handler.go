package handler

import (
	"context"
	"douyin-project/common/token"
	"douyin-project/service/user/logic"
	"douyin-project/service/user/svcctx"
	"douyin-project/service/user/types"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

// UserAction no practical effect, just check if token is valid
func UserRegisterHandler(serviceCtx *svcctx.ServiceContext) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		req := &types.UserRegisterReq{}
		if err := c.BindAndValidate(req); err != nil {
			// c.JSON(http.StatusOK, types.UserRegisterResp{
			// 	StatusCode: -1,
			// 	StatusMsg:  "parse register req fail",
			// })
			return
		}

		l := logic.NewUserLogic(serviceCtx)
		// resp, err := l.UserRegister(ctx, &types.UserRegisterLogicReq{
		// 	Name:     req.Name,
		// 	Password: req.Password,
		// })
		l.UserRegister(ctx, &types.UserRegisterLogicReq{
			Name:     req.Name,
			Password: req.Password,
		})

		//由于下面还会调用login，所以这里就没有信息了

		// if err != nil {
		// 	c.JSON(http.StatusOK, types.UserRegisterResp{
		// 		StatusCode: http.StatusBadRequest,
		// 		StatusMsg:  "register fail",
		// 	})
		// } else {
		// 	c.JSON(http.StatusOK, types.UserRegisterResp{
		// 		StatusCode: 0,
		// 		StatusMsg:  "success",
		// 		UserId:     resp.UserId,
		// 	})
		// }
	}
}

// UserAction no practical effect, just check if token is valid
func UserInfoHandler(serviceCtx *svcctx.ServiceContext) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		u, _ := c.Get(serviceCtx.IdentityKey)
		tokenUser := u.(*token.TokenUser)

		req := &types.UserInfoReq{}
		if err := c.BindAndValidate(req); err != nil {
			return
		}

		l := logic.NewUserLogic(serviceCtx)
		// resp, err := l.UserRegister(ctx, &types.UserRegisterLogicReq{
		// 	Name:     req.Name,
		// 	Password: req.Password,
		// })
		resp, err := l.UserInfo(ctx, &types.UserInfoLogicReq{
			UserId:        req.UserId,
			CurrentUserId: tokenUser.Id,
		})
		if err != nil {
			c.JSON(http.StatusOK, types.UserInfoResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
