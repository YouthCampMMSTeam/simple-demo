package main

import (
	"context"
	"douyin-project/microservice/favorite/model"
	"douyin-project/microservice/favorite/rpc/config"
	"douyin-project/microservice/favorite/rpc/kitex_gen/favorite"
	"douyin-project/microservice/favorite/rpc/svcctx"

	"github.com/jinzhu/copier"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct {
	serviceCtx *svcctx.ServiceContext
}

func NewFavoriteServiceImpl(c *config.Config) *FavoriteServiceImpl {
	return &FavoriteServiceImpl{
		serviceCtx: svcctx.NewServiceContext(c),
	}
}

// FindByVideoIdAndUserId implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FindByVideoIdAndUserId(ctx context.Context, req *favorite.FindByVideoIdAndUserIdRequest) (resp *favorite.FindByVideoIdAndUserIdResp, err error) {
	results, err := s.serviceCtx.FavoriteModel.FindByVideoIdAndUserId(ctx, req.VideoId, req.UserId)
	if err != nil {
		return nil, err
	}
	resp = &favorite.FindByVideoIdAndUserIdResp{}
	copier.Copy(&resp.FavoriteList, &results)
	return resp, nil
}

// FindByUserId implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FindByUserId(ctx context.Context, req *favorite.FindByUserIdRequest) (resp *favorite.FindByUserIdResp, err error) {
	results, err := s.serviceCtx.FavoriteModel.FindByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	resp = &favorite.FindByUserIdResp{}
	copier.Copy(&resp.FavoriteList, &results)
	return resp, nil
}

// Insert implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Insert(ctx context.Context, req *favorite.InsertRequest) (resp *favorite.InsertResp, err error) {
	var f model.Favorite
	copier.Copy(&f, req.Favorite)
	if s.serviceCtx.FavoriteModel.Insert(ctx, &f); err != nil {
		return nil, err
	}
	resp = &favorite.InsertResp{
		FavoriteId: int64(f.ID),
	}
	return resp, nil
}

// Delete implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Delete(ctx context.Context, req *favorite.DeleteRequest) (resp *favorite.DeleteResp, err error) {
	if err := s.serviceCtx.FavoriteModel.Delete(ctx, req.FavoriteId); err != nil {
		return nil, err
	}
	resp = favorite.NewDeleteResp()
	return resp, nil
}
