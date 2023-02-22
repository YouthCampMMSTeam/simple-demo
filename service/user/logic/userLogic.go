package logic

import (
	"context"
	"douyin-project/microservice/user/rpc/kitex_gen/user"
	"douyin-project/service/user/svcctx"
	"douyin-project/service/user/types"
	"errors"
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
