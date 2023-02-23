package handler

import (
	"context"
	"douyin-project/common/token"
	"douyin-project/service/publish/logic"
	"douyin-project/service/publish/oss"
	"douyin-project/service/publish/svcctx"
	"douyin-project/service/publish/types"
	"log"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

// PublishAction no practical effect, just check if token is valid
func PublishActionHandler(serviceCtx *svcctx.ServiceContext) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		u, _ := c.Get(serviceCtx.IdentityKey)
		//值得注意的是，这里的TokenUser必须和login那边是完全一致的，所以这里把定义放到common中间
		//如果仅仅是结构体中间字段内容相同是不行的
		tokenUser := u.(*token.TokenUser)

		publishActionReq := &types.PublishActionReq{}
		if err := c.BindAndValidate(publishActionReq); err != nil {
			c.JSON(http.StatusOK, types.PublishActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		data, err := c.FormFile("data")

		//TODO 补一下从[]data转到url的连接
		uuidNum := serviceCtx.UuidGenerator.Generate().Int64()
		url, err := oss.TranToUrl(data, strconv.FormatInt(uuidNum, 10))
		if err != nil {
			log.Fatalf("url: %+v\n", err)
		}

		publishActionLogicReq := &types.PublishActionLogicReq{
			CurrentUserId: tokenUser.Id,
			PlayUrl:       url,
			CoverUrl:      url,
			Title:         publishActionReq.Title,
		}

		l := logic.NewPublishLogic(serviceCtx)
		_, err = l.PublishAction(ctx, publishActionLogicReq)
		if err != nil {
			c.JSON(http.StatusOK, types.PublishActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		c.JSON(http.StatusOK, types.PublishActionResp{
			StatusMsg:  "success",
			StatusCode: 0,
		})

	}
}
func PublishListHandler(serviceCtx *svcctx.ServiceContext) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		u, _ := c.Get(serviceCtx.IdentityKey)
		tokenUser := u.(*token.TokenUser)

		publishListReq := &types.PublishListReq{}
		if err := c.BindAndValidate(publishListReq); err != nil {
			c.JSON(http.StatusOK, types.PublishActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		l := logic.NewPublishLogic(serviceCtx)
		resp, err := l.PublishList(ctx, &types.PublishListLogicReq{
			CurrentUserId: tokenUser.Id,
			UserId:        publishListReq.UserId,
		})
		if err != nil {
			c.JSON(http.StatusOK, types.PublishActionResp{
				StatusMsg:  err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		c.JSON(http.StatusOK, resp)

	}
}
