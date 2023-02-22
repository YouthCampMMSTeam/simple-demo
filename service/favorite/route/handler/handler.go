package handler

import (
	"context"
	"douyin-project/common/token"
	"douyin-project/service/favorite/logic"
	"douyin-project/service/favorite/svcctx"
	"douyin-project/service/favorite/types"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteActionHandler(serviceCtx *svcctx.ServiceContext) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		u, _ := c.Get(serviceCtx.IdentityKey)

		//值得注意的是，这里的TokenUser必须和login那边是完全一致的，所以这里把定义放到common中间
		//如果仅仅是结构体中间字段内容相同是不行的
		tokenUser := u.(*token.TokenUser)

		favoriteActionReq := &types.FavoriteActionReq{}
		if err := c.BindAndValidate(favoriteActionReq); err != nil {
			c.JSON(http.StatusOK, types.FavoriteActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		l := logic.NewFavoriteLogic(serviceCtx)
		_, err := l.FavoriteAction(ctx, &types.FavoriteActionLogicReq{
			UserId:     tokenUser.Id,
			VideoId:    favoriteActionReq.VideoId,
			ActionType: favoriteActionReq.ActionType,
		})
		if err != nil {
			c.JSON(http.StatusOK, types.FavoriteActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		c.JSON(http.StatusOK, types.FavoriteActionResp{
			StatusMsg:  "success",
			StatusCode: 0,
		})

	}
}
func FavoriteListHandler(serviceCtx *svcctx.ServiceContext) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		u, _ := c.Get(serviceCtx.IdentityKey)
		tokenUser := u.(*token.TokenUser)

		favoriteListReq := &types.FavoriteListReq{}
		if err := c.BindAndValidate(favoriteListReq); err != nil {
			c.JSON(http.StatusOK, types.FavoriteActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		l := logic.NewFavoriteLogic(serviceCtx)
		resp, err := l.FavoriteList(ctx, &types.FavoriteListLogicReq{
			CurrentUserId: tokenUser.Id,
			UserId:        favoriteListReq.UserId,
		})
		if err != nil {
			c.JSON(http.StatusOK, types.FavoriteActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		c.JSON(http.StatusOK, resp)

	}
}
