package handler

import (
	"context"
	"douyin-project/service/user/logic"
	"douyin-project/service/user/svcctx"
	"douyin-project/service/user/types"

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

// if _, exist := usersLoginInfo[token]; exist {
// 	c.JSON(http.StatusOK, Response{StatusCode: 0})
// } else {
// 	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
// }
