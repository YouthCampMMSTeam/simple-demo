package logic

import (
	"context"
	"douyin-project/microservice/relation/rpc/kitex_gen/relation"
	"douyin-project/microservice/user/rpc/kitex_gen/user"
	"douyin-project/service/user/svcctx"
	"douyin-project/service/user/types"
	"errors"

	"github.com/jinzhu/copier"
)

type UserLogic struct {
	serviceCtx *svcctx.ServiceContext
}

func NewUserLogic(svcCtx *svcctx.ServiceContext) *UserLogic {
	return &UserLogic{
		serviceCtx: svcCtx,
	}
}

func (l *UserLogic) UserLogin(ctx context.Context, req *types.UserLoginLogicReq) (*types.UserLoginLogicResp, error) {
	//还是得进行幂等判断
	findByNameResp, err := l.serviceCtx.UserRpc.FindByName(ctx, &user.FindByNameRequest{
		UserName: req.Name,
	})

	if err != nil {
		return nil, err
	}
	if len(findByNameResp.UserList) == 0 {
		return nil, errors.New("user not exists")
	}
	if findByNameResp.UserList[0].Password != req.Password {
		return nil, errors.New("wrong password")
	}

	return &types.UserLoginLogicResp{
		UserId: findByNameResp.UserList[0].Id,
	}, nil
}

func (l *UserLogic) UserRegister(ctx context.Context, req *types.UserRegisterLogicReq) (*types.UserRegisterLogicResp, error) {

	//还是得进行幂等判断
	findByNameResp, err := l.serviceCtx.UserRpc.FindByName(ctx, &user.FindByNameRequest{
		UserName: req.Name,
	})
	if err != nil {
		return nil, err
	}
	if len(findByNameResp.UserList) != 0 {
		return nil, errors.New("user already exists")
	}

	insertResp, err := l.serviceCtx.UserRpc.Insert(ctx, &user.InsertRequest{
		User: &user.User{
			Name:          req.Name,
			Password:      req.Password,
			FollowCount:   0,
			FollowerCount: 0,
		},
	})
	if err != nil {
		return nil, err
	}

	return &types.UserRegisterLogicResp{
		UserId: insertResp.UserId,
	}, nil
}

func (l *UserLogic) UserInfo(ctx context.Context, req *types.UserInfoLogicReq) (*types.UserInfoResp, error) {

	//首先判断token对应用户是否存在
	findByUserIdResp, err := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{
		UserId: req.CurrentUserId,
	})
	if err != nil {
		return nil, err
	} else if len(findByUserIdResp.UserList) == 0 {
		return nil, errors.New("current user not exists")
	}

	//查询目标用户信息
	findByTargetUserIdResp, err := l.serviceCtx.UserRpc.FindByUserId(ctx, &user.FindByUserIdRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	} else if len(findByTargetUserIdResp.UserList) == 0 {
		return nil, errors.New("target user not exists")
	}

	resp := &types.UserInfoResp{}
	copier.Copy(&resp.User, findByTargetUserIdResp.UserList[0])

	//判断是否有相互关注
	selectRelationResp, err := l.serviceCtx.RelationRpc.SelectRelation(ctx, &relation.SelectRelationRequest{
		FollowId:   req.UserId,
		FollowerId: req.CurrentUserId,
	})
	if err != nil {
		return nil, err
	}
	if len(selectRelationResp.RelationList) == 0 {
		resp.User.IsFollow = false
	} else {
		resp.User.IsFollow = true
	}

	resp.StatusCode = 0
	resp.StatusMsg = "success"
	return resp, nil
}
