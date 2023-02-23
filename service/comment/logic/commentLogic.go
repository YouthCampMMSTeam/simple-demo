package logic

import (
	"context"
	"douyin-project/microservice/comment/rpc/kitex_gen/comment"
	"douyin-project/microservice/relation/rpc/kitex_gen/relation"
	"douyin-project/microservice/user/rpc/kitex_gen/user"
	"douyin-project/microservice/video/rpc/kitex_gen/video"
	"douyin-project/service/comment/svcctx"
	"douyin-project/service/comment/types"
	"errors"

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

//评论操作
func (l *CommentLogic) CommentAction(ctx context.Context, req *types.CommentActionLogicReq) (*types.CommentActionLogicResp, error) {

	//首先判断token对应用户是否存在
	findByUserIdResp, err := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{
		UserId: req.CurrentUserId,
	})
	if err != nil {
		return nil, err
	} else if len(findByUserIdResp.UserList) == 0 {
		return nil, errors.New("current user not exists")
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

	resp := &types.CommentActionLogicResp{}
	//然后判断一下是发布评论还是删除评论
	if req.ActionType == 1 {
		//发布评论
		insertRequest := &comment.InsertReq{
			Comment: &comment.Comment{
				VideoId: req.VideoId,
				UserId:  req.CurrentUserId,
				Content: req.CommentText,
			},
		}
		insertResp, err := l.serviceCtx.CommentRpc.Insert(ctx, insertRequest)
		if err != nil {
			return nil, err
		}
		resp.Comment.CreateDate = insertResp.CreateDate
		copier.Copy(&resp.Comment, insertRequest)
		resp.Comment.User = &types.User{}
		copier.Copy(resp.Comment.User, findByUserIdResp.UserList[0])
	} else if req.ActionType == 2 {
		//删除评论
		if _, err := l.serviceCtx.CommentRpc.Delete(ctx, &comment.DeleteReq{
			CommentId: req.CommentId,
		}); err != nil {
			return nil, err
		}
	}

	return resp, nil
}

//用户点赞列表
func (l *CommentLogic) CommentList(ctx context.Context, req *types.CommentListLogicReq) (*types.CommentListResp, error) {

	//首先判断token对应用户是否存在
	findUserIdResp, err := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{
		UserId: req.CurrentUserId,
	})
	if err != nil {
		return nil, err
	} else if len(findUserIdResp.UserList) == 0 {
		return nil, errors.New("user not exists")
	}

	//查询该视频下面的每一个评论
	findByVideoIdResp, err := l.serviceCtx.CommentRpc.FindByVideoId(ctx, &comment.FindByVideoIdReq{
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}

	//对每一个评论的用户进行分析
	var commentDetails []*types.Comment
	for _, com := range findByVideoIdResp.CommentList {
		//对于每个评论创建对象
		c := &types.Comment{
			User: &types.User{}, //因为是指针，别忘了再给一个对象
		}
		copier.Copy(c, com)

		//剩下确认一下作者信息

		//获取评论的用户信息
		findByUserIdResp, _ := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{UserId: com.UserId})
		copier.Copy(c.User, findByUserIdResp.UserList[0])

		//判断评论对应用户和当前用户关系
		selectRelationResp, _ := l.serviceCtx.RelationRpc.SelectRelation(ctx, &relation.SelectRelationRequest{
			FollowId:   findByUserIdResp.UserList[0].Id,
			FollowerId: req.CurrentUserId,
		})
		if len(selectRelationResp.RelationList) == 0 {
			//没关注
			c.User.IsFollow = false
		} else {
			//关注了
			c.User.IsFollow = true
		}
		commentDetails = append(commentDetails, c)
	}

	//TODO 包的有点乱，不过我是觉得这样不用多次复制，效率高点，但获取使用指针然后包多层是一个更好的选择
	resp := &types.CommentListResp{
		CommentList: commentDetails,
		StatusCode:  0,
		StatusMsg:   "success",
	}
	return resp, nil
}
