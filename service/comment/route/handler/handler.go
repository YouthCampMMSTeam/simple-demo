package handler

import (
	"context"
	"douyin-project/common/token"
	"douyin-project/service/comment/logic"
	"douyin-project/service/comment/svcctx"
	"douyin-project/service/comment/types"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
)

// CommentAction no practical effect, just check if token is valid
func CommentActionHandler(serviceCtx *svcctx.ServiceContext) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		u, _ := c.Get(serviceCtx.IdentityKey)
		//值得注意的是，这里的TokenUser必须和login那边是完全一致的，所以这里把定义放到common中间
		//如果仅仅是结构体中间字段内容相同是不行的
		tokenUser := u.(*token.TokenUser)

		commentActionReq := &types.CommentActionReq{}
		if err := c.BindAndValidate(commentActionReq); err != nil {
			c.JSON(http.StatusOK, types.CommentActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		commentActionLogicReq := &types.CommentActionLogicReq{}
		copier.Copy(commentActionLogicReq, commentActionReq)
		commentActionLogicReq.CurrentUserId = tokenUser.Id

		l := logic.NewCommentLogic(serviceCtx)
		commentActionLogicResp, err := l.CommentAction(ctx, commentActionLogicReq)
		if err != nil {
			c.JSON(http.StatusOK, types.CommentActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		c.JSON(http.StatusOK, types.CommentActionResp{
			StatusMsg:  "success",
			StatusCode: 0,
			Comment:    commentActionLogicResp.Comment,
		})

	}
}
func CommentListHandler(serviceCtx *svcctx.ServiceContext) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		u, _ := c.Get(serviceCtx.IdentityKey)
		tokenUser := u.(*token.TokenUser)

		commentListReq := &types.CommentListReq{}
		if err := c.BindAndValidate(commentListReq); err != nil {
			c.JSON(http.StatusOK, types.CommentActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		l := logic.NewCommentLogic(serviceCtx)
		resp, err := l.CommentList(ctx, &types.CommentListLogicReq{
			CurrentUserId: tokenUser.Id,
			VideoId:       commentListReq.VideoId,
		})
		if err != nil {
			c.JSON(http.StatusOK, types.CommentActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		c.JSON(http.StatusOK, resp)

	}
}
