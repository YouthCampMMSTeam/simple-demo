package main

import (
	"context"
	"douyin-project/microservice/video/model"
	"douyin-project/microservice/video/rpc/config"
	"douyin-project/microservice/video/rpc/kitex_gen/video"
	"douyin-project/microservice/video/rpc/svcctx"

	"github.com/jinzhu/copier"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct {
	serviceCtx svcctx.ServiceContext
}

func NewVideoServiceImpl(c *config.Config) *VideoServiceImpl {
	return &VideoServiceImpl{
		serviceCtx: *svcctx.NewServiceContext(c),
	}
}

// FindOrderByTime implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FindOrderByTime(ctx context.Context, req *video.FindOrderByTimeReq) (resp *video.FindOrderByTimeResp, err error) {
	results, err := s.serviceCtx.VideoModel.FindOrderByTime(ctx, req.LimitNum)
	if err != nil {
		return nil, err
	}
	resp = &video.FindOrderByTimeResp{}
	copier.Copy(&resp.VideoList, &results)
	return resp, nil
}

// FindWithTimeLimit implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FindWithTimeLimit(ctx context.Context, req *video.FindWithTimeLimitReq) (resp *video.FindWithTimeLimitResp, err error) {
	results, err := s.serviceCtx.VideoModel.FindWithTimeLimit(ctx, req.LatestTime)
	if err != nil {
		return nil, err
	}
	resp = &video.FindWithTimeLimitResp{}
	copier.Copy(&resp.VideoList, &results)
	if len(results) > 0 {
		resp.NextTime = results[len(results)-1].CreatedAt.Unix()
	}
	return resp, nil
}

// FindByVideoId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FindByVideoId(ctx context.Context, req *video.FindByVideoIdReq) (resp *video.FindByVideoIdResp, err error) {
	results, err := s.serviceCtx.VideoModel.FindByVideoId(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	resp = &video.FindByVideoIdResp{}
	copier.Copy(&resp.VideoList, &results)
	return resp, nil
}

// FindByUserId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FindByUserId(ctx context.Context, req *video.FindByUserIdReq) (resp *video.FindByUserIdResp, err error) {
	results, err := s.serviceCtx.VideoModel.FindByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	resp = &video.FindByUserIdResp{}
	copier.Copy(&resp.VideoList, &results)
	return resp, nil
}

// Insert implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Insert(ctx context.Context, req *video.InsertReq) (resp *video.InsertResp, err error) {
	var f model.Video
	copier.Copy(&f, req.Video)
	if s.serviceCtx.VideoModel.Insert(ctx, &f); err != nil {
		return nil, err
	}
	resp = &video.InsertResp{
		VideoId: int64(f.ID),
	}
	return resp, nil
}

// Update implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Update(ctx context.Context, req *video.UpdateReq) (resp *video.UpdateResp, err error) {
	u := &model.Video{}
	copier.Copy(u, req.Video)
	//ID对应Id转不过来，手动转一下
	u.ID = uint(req.Video.Id)
	if err := s.serviceCtx.VideoModel.Update(ctx, u); err != nil {
		return nil, err
	}
	return resp, nil
}

// FavoriteCountModified implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteCountModified(ctx context.Context, req *video.FavoriteCountModifiedReq) (resp *video.FavoriteCountModifiedResp, err error) {
	if err := s.serviceCtx.VideoModel.FavoriteCountModified(ctx, uint(req.VideoId), req.PosOrNeg); err != nil {
		return nil, err
	}
	return resp, nil
}

// CommentCountModified implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentCountModified(ctx context.Context, req *video.CommentCountModifiedReq) (resp *video.CommentCountModifiedResp, err error) {
	if err := s.serviceCtx.VideoModel.CommentCountModified(ctx, uint(req.VideoId), req.PosOrNeg); err != nil {
		return nil, err
	}
	return resp, nil
}
