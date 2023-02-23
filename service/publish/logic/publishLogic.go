package logic

import (
	"context"
	"douyin-project/microservice/favorite/rpc/kitex_gen/favorite"
	"douyin-project/microservice/relation/rpc/kitex_gen/relation"
	"douyin-project/microservice/user/rpc/kitex_gen/user"
	"douyin-project/microservice/video/rpc/kitex_gen/video"
	"douyin-project/service/publish/svcctx"
	"douyin-project/service/publish/types"
	"errors"

	"github.com/jinzhu/copier"
)

type PublishLogic struct {
	serviceCtx *svcctx.ServiceContext
}

func NewPublishLogic(svcCtx *svcctx.ServiceContext) *PublishLogic {
	return &PublishLogic{
		serviceCtx: svcCtx,
	}
}

//发布视频
func (l *PublishLogic) PublishAction(ctx context.Context, req *types.PublishActionLogicReq) (*types.PublishActionLogicResp, error) {

	//首先判断token对应用户是否存在
	findByUserIdResp, err := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{
		UserId: req.CurrentUserId,
	})
	if err != nil {
		return nil, err
	} else if len(findByUserIdResp.UserList) == 0 {
		return nil, errors.New("current user not exists")
	}

	//创建一下新的video
	v := &video.Video{
		PlayUrl:       req.PlayUrl,
		CoverUrl:      req.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		AuthorId:      req.CurrentUserId,
	}
	_, err = l.serviceCtx.VideoRpc.Insert(ctx, &video.InsertReq{
		Video: v,
	})
	if err != nil {
		return nil, err
	}

	return &types.PublishActionLogicResp{}, nil
}

//用户的视频发布列表
func (l *PublishLogic) PublishList(ctx context.Context, req *types.PublishListLogicReq) (*types.PublishListResp, error) {

	//首先判断token对应用户是否存在
	findUserIdResp, err := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{
		UserId: req.CurrentUserId,
	})
	if err != nil {
		return nil, err
	} else if len(findUserIdResp.UserList) == 0 {
		return nil, errors.New("user not exists")
	}

	//查询目标用户是否存在
	findTargetUserIdResp, err := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{
		UserId: req.CurrentUserId,
	})
	if err != nil {
		return nil, err
	} else if len(findTargetUserIdResp.UserList) == 0 {
		return nil, errors.New("target user not exists")
	}

	//查看目标用户对应的所有视频
	findByVideoIdResp, err := l.serviceCtx.VideoRpc.FindByUserId(ctx, &video.FindByUserIdReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	selectRelationReq, err := l.serviceCtx.RelationRpc.SelectRelation(ctx, &relation.SelectRelationRequest{
		FollowId:   req.UserId,
		FollowerId: req.CurrentUserId,
	})
	if err != nil {
		return nil, err
	}

	var isFollow bool
	if len(selectRelationReq.RelationList) == 0 {
		//没有关注
		isFollow = false
	} else {
		isFollow = true
	}

	//对每一个视频进行具体信息获取
	var videoDetails []*types.Video
	for _, video := range findByVideoIdResp.VideoList {
		//对于每个评论创建对象
		v := &types.Video{
			Author: &types.User{}, //因为是指针，别忘了再给一个对象
		}
		copier.Copy(v, video)

		//前面统一查询目标用户信息了
		copier.Copy(v.Author, findTargetUserIdResp.UserList[0])

		v.Author.IsFollow = isFollow
		//最后是token的用户是否给当前查询用户的这个视频点赞了
		findByVideoIdAndUserIdResp, _ := l.serviceCtx.FavoriteRpc.FindByVideoIdAndUserId(ctx, &favorite.FindByVideoIdAndUserIdRequest{
			UserId:  req.CurrentUserId,
			VideoId: video.AuthorId,
		})
		if len(findByVideoIdAndUserIdResp.FavoriteList) == 0 {
			//没点赞
			v.IsFavorite = false
		} else {
			//点赞了
			v.IsFavorite = true
		}
		videoDetails = append(videoDetails, v)
	}

	//TODO 包的有点乱，不过我是觉得这样不用多次复制，效率高点，但获取使用指针然后包多层是一个更好的选择
	resp := &types.PublishListResp{
		VideoList:  videoDetails,
		StatusCode: 0,
		StatusMsg:  "success",
	}
	return resp, nil
}
