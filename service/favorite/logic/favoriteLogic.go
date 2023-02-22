package logic

import (
	"context"
	"douyin-project/microservice/favorite/rpc/kitex_gen/favorite"
	"douyin-project/microservice/relation/rpc/kitex_gen/relation"
	"douyin-project/microservice/user/rpc/kitex_gen/user"
	"douyin-project/microservice/video/rpc/kitex_gen/video"
	"douyin-project/service/favorite/svcctx"
	"douyin-project/service/favorite/types"
	"errors"

	"github.com/jinzhu/copier"
)

type FavoriteLogic struct {
	serviceCtx *svcctx.ServiceContext
}

func NewFavoriteLogic(svcCtx *svcctx.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		serviceCtx: svcCtx,
	}
}

//点赞操作
func (l *FavoriteLogic) FavoriteAction(ctx context.Context, req *types.FavoriteActionLogicReq) (*types.FavoriteActionLogicResp, error) {

	//首先判断是否存在该用户
	findUserIdResp, err := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	} else if len(findUserIdResp.UserList) == 0 {
		return nil, errors.New("user not exists")
	}

	//然后判断是否存在该视频
	FindByVideoIdResp, err := l.serviceCtx.VideoRpc.FindByVideoId(ctx, &video.FindByVideoIdReq{
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	} else if len(FindByVideoIdResp.VideoList) == 0 {
		return nil, errors.New("video not exists")
	}

	//此时查看是否有查找点赞信息
	findByVideoIdAndUserIdResp, err := l.serviceCtx.FavoriteRpc.FindByVideoIdAndUserId(ctx, &favorite.FindByVideoIdAndUserIdRequest{
		VideoId: req.VideoId,
		UserId:  req.UserId,
	})
	if err != nil {
		return nil, err
	}

	//进行一些幂等性判断
	if req.ActionType == 1 && len(findByVideoIdAndUserIdResp.FavoriteList) == 0 {
		//要点赞，但点赞信息不存在，即为新的点赞操作
		_, err := l.serviceCtx.FavoriteRpc.Insert(ctx, &favorite.InsertRequest{
			Favorite: &favorite.Favorite{
				VideoId: req.VideoId,
				UserId:  req.UserId,
			},
		})
		if err != nil {
			return nil, err
		}
		//修改视频点赞信息
		if _, err = l.serviceCtx.VideoRpc.FavoriteCountModified(ctx, &video.FavoriteCountModifiedReq{
			VideoId:  req.VideoId,
			PosOrNeg: true,
		}); err != nil {
			return nil, err
		}

	} else if req.ActionType == 2 && len(findByVideoIdAndUserIdResp.FavoriteList) == 1 {
		// 取消点赞，而点赞信息存在， 即为取消点赞操作
		_, err := l.serviceCtx.FavoriteRpc.Delete(ctx, &favorite.DeleteRequest{
			FavoriteId: findByVideoIdAndUserIdResp.FavoriteList[0].Id,
		})
		if err != nil {
			return nil, err
		}

		//修改视频点赞信息
		if _, err = l.serviceCtx.VideoRpc.FavoriteCountModified(ctx, &video.FavoriteCountModifiedReq{
			VideoId:  req.VideoId,
			PosOrNeg: false,
		}); err != nil {
			return nil, err
		}
	}
	return &types.FavoriteActionLogicResp{}, nil
}

//用户点赞列表
func (l *FavoriteLogic) FavoriteList(ctx context.Context, req *types.FavoriteListLogicReq) (*types.FavoriteListResp, error) {

	//首先判断是否存在该用户
	findUserIdResp, err := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{
		UserId: req.CurrentUserId,
	})
	if err != nil {
		return nil, err
	} else if len(findUserIdResp.UserList) == 0 {
		return nil, errors.New("user not exists")
	}

	//查询目标用户喜欢的视频id
	findByUserIdResp, err := l.serviceCtx.FavoriteRpc.FindByUserId(ctx, &favorite.FindByUserIdRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	//对每一个视频进行分析
	var videosDetails []*types.Video
	for _, f := range findByUserIdResp.FavoriteList {
		v := &types.Video{}
		//TODO 下面部分err没判断
		//获取视频信息
		findByVideoIdResp, _ := l.serviceCtx.VideoRpc.FindByVideoId(ctx, &video.FindByVideoIdReq{VideoId: f.VideoId})
		copier.Copy(v, findByVideoIdResp.VideoList[0])
		v.IsFavorite = true

		//获取作者信息
		findByUserIdResp, _ := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{UserId: findByVideoIdResp.VideoList[0].AuthorId})
		copier.Copy(v.Author, findByUserIdResp.UserList[0])

		//判断作者和当前用户关系
		selectRelationResp, _ := l.serviceCtx.RelationRpc.SelectRelation(ctx, &relation.SelectRelationRequest{
			FollowId:   findByUserIdResp.UserList[0].Id,
			FollowerId: req.CurrentUserId,
		})
		if len(selectRelationResp.RelationList) == 0 {
			//没关注
			v.Author.IsFollow = false
		} else {
			//关注了
			v.Author.IsFollow = true
		}
		videosDetails = append(videosDetails, v)
	}

	resp := &types.FavoriteListResp{
		VideoList:  videosDetails,
		StatusCode: 0,
		StatusMsg:  "success",
	}
	return resp, nil
}
