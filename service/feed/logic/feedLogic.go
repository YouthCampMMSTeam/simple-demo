package logic

import (
	"context"
	"douyin-project/microservice/favorite/rpc/kitex_gen/favorite"
	"douyin-project/microservice/relation/rpc/kitex_gen/relation"
	"douyin-project/microservice/user/rpc/kitex_gen/user"
	"douyin-project/microservice/video/rpc/kitex_gen/video"
	"douyin-project/service/feed/svcctx"
	"douyin-project/service/feed/types"

	"github.com/jinzhu/copier"
)

type CommentLogic struct {
	serviceCtx *svcctx.ServiceContext
}

func NewCommentLogic(svcCtx *svcctx.ServiceContext) *CommentLogic {
	return &CommentLogic{
		serviceCtx: svcCtx,
	}
}

func (l *CommentLogic) Feed(ctx context.Context, req *types.FeedLogicReq) (*types.FeedResp, error) {
	//直接按照时间顺序获取视频就可以了
	findWithTimeLimitResp, err := l.serviceCtx.VideoRpc.FindWithTimeLimit(ctx, &video.FindWithTimeLimitReq{
		LatestTime: req.LatestTime,
	})
	if err != nil {
		return nil, err
	}
	var videoDetails []*types.Video
	for _, v := range findWithTimeLimitResp.VideoList {
		videoDetail := &types.Video{
			Author: &types.User{},
		}
		copier.Copy(videoDetail, v)

		//获取该视频的作者信息
		findByUserIdRequest, _ := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{
			UserId: v.AuthorId,
		})
		copier.Copy(videoDetail.Author, findByUserIdRequest.UserList[0])

		//如果有当前用户，则需要判断中间和当前用户是否有互关，以及和对应的视频是否有点赞
		if req.CurrentUserId != 0 {
			//判断互关
			selectRelationResp, _ := l.serviceCtx.RelationRpc.SelectRelation(ctx, &relation.SelectRelationRequest{
				FollowId:   v.AuthorId,
				FollowerId: req.CurrentUserId,
			})
			if len(selectRelationResp.RelationList) == 0 {
				//没有关注
				videoDetail.Author.IsFollow = false
			} else {
				videoDetail.Author.IsFollow = true
			}

			//判断点赞
			findByVideoIdAndUserIdResp, _ := l.serviceCtx.FavoriteRpc.FindByVideoIdAndUserId(ctx, &favorite.FindByVideoIdAndUserIdRequest{
				VideoId: v.Id,
				UserId:  req.CurrentUserId,
			})
			if len(findByVideoIdAndUserIdResp.FavoriteList) == 0 {
				//没有关注
				videoDetail.IsFavorite = false
			} else {
				videoDetail.IsFavorite = true
			}
		}
		videoDetails = append(videoDetails, videoDetail)
	}

	resp := &types.FeedResp{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videoDetails,
		NextTime:   findWithTimeLimitResp.NextTime,
	}

	return resp, nil
}
