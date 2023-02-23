package handler

import (
	"context"
	"douyin-project/common/token"
	"douyin-project/service/feed/logic"
	"douyin-project/service/feed/svcctx"
	"douyin-project/service/feed/types"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

// CommentAction no practical effect, just check if token is valid
func FeedHandler(serviceCtx *svcctx.ServiceContext) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		u, _ := c.Get(serviceCtx.IdentityKey)
		//值得注意的是，这里的TokenUser必须和login那边是完全一致的，所以这里把定义放到common中间
		//如果仅仅是结构体中间字段内容相同是不行的
		tokenUser := u.(*token.TokenUser)

		feedReq := &types.FeedReq{}
		if err := c.BindAndValidate(feedReq); err != nil {
			c.JSON(http.StatusOK, types.FeedResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		l := logic.NewCommentLogic(serviceCtx)
		feedResp, err := l.Feed(ctx, &types.FeedLogicReq{
			CurrentUserId: tokenUser.Id,
			LatestTime:    feedReq.LatestTime,
		})

		if err != nil {
			c.JSON(http.StatusOK, types.FeedResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		c.JSON(http.StatusOK, feedResp)

	}
}
