package main

import (
	"context"
	"douyin-project/microservice/user/model"
	"douyin-project/microservice/user/rpc/config"
	user "douyin-project/microservice/user/rpc/kitex_gen/user"
	"douyin-project/microservice/user/rpc/svcctx"
	"fmt"

	"github.com/jinzhu/copier"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	serviceCtx svcctx.ServiceContext
}

func NewUserServiceImpl(c *config.Config) *UserServiceImpl {
	return &UserServiceImpl{
		serviceCtx: *svcctx.NewServiceContext(c),
	}
}

// FindByName implements the UserServiceImpl interface.
func (s *UserServiceImpl) FindByName(ctx context.Context, req *user.FindByNameRequest) (resp *user.FindByNameResp, err error) {
	results, err := s.serviceCtx.UserModel.FindByName(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	resp = &user.FindByNameResp{}
	copier.Copy(&resp.UserList, &results)
	return resp, nil
}

// FindByUserId implements the UserServiceImpl interface.
func (s *UserServiceImpl) FindByUserId(ctx context.Context, req *user.FindByUserIdRequest) (resp *user.FindByUserIdResp, err error) {
	results, err := s.serviceCtx.UserModel.FindByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	resp = &user.FindByUserIdResp{}
	copier.Copy(&resp.UserList, &results)
	return resp, nil
}

// Insert implements the UserServiceImpl interface.
func (s *UserServiceImpl) Insert(ctx context.Context, req *user.InsertRequest) (resp *user.InsertResp, err error) {
	var u model.User
	copier.Copy(&u, req.User)

	if err := s.serviceCtx.UserModel.Insert(ctx, &u); err != nil {
		return nil, err
	}
	resp = &user.InsertResp{
		UserId: int64(u.ID),
	}
	return resp, nil
}

// Update implements the UserServiceImpl interface.
func (s *UserServiceImpl) Update(ctx context.Context, req *user.UpdateRequest) (resp *user.UpdateResp, err error) {
	// TODO: Your code here...
	u := &model.User{}
	copier.Copy(u, req.User)
	fmt.Printf("u: %+v\n", u)
	fmt.Printf("req.User: %+v\n", req.User)
	if err := s.serviceCtx.UserModel.Update(ctx, u); err != nil {
		return nil, err
	}
	return resp, nil
}
