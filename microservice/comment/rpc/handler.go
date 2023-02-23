package main

import (
	"context"
	"douyin-project/microservice/comment/model"
	"douyin-project/microservice/comment/rpc/config"
	"douyin-project/microservice/comment/rpc/kitex_gen/comment"
	"douyin-project/microservice/comment/rpc/svcctx"

	"github.com/jinzhu/copier"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct {
	serviceCtx *svcctx.ServiceContext
}

func NewCommentServiceImpl(c *config.Config) *CommentServiceImpl {
	return &CommentServiceImpl{
		serviceCtx: svcctx.NewServiceContext(c),
	}
}

// FindCommentByVideoIdLimit30 implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) FindCommentByVideoIdLimit30(ctx context.Context, req *comment.FindCommentByVideoIdLimit30Req) (resp *comment.FindCommentByVideoIdLimit30Resp, err error) {
	results, err := s.serviceCtx.CommentModel.FindByVideoIdLimit30(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	resp = &comment.FindCommentByVideoIdLimit30Resp{}
	copier.Copy(&resp.CommentList, &results)
	var MMDD = "01/02" //MMDD日期格式
	for ind, r := range results {
		resp.CommentList[ind].CreateDate = r.CreatedAt.Format(MMDD)
	}
	return resp, nil
}

// FindByVideoId implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) FindByVideoId(ctx context.Context, req *comment.FindByVideoIdReq) (resp *comment.FindByVideoIdResp, err error) {
	results, err := s.serviceCtx.CommentModel.FindByVideoId(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	resp = &comment.FindByVideoIdResp{}
	copier.Copy(&resp.CommentList, &results)
	var MMDD = "01/02" //MMDD日期格式
	for ind, r := range results {
		resp.CommentList[ind].CreateDate = r.CreatedAt.Format(MMDD)
	}
	return resp, nil
}

// Insert implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Insert(ctx context.Context, req *comment.InsertReq) (resp *comment.InsertResp, err error) {
	var c model.Comment
	copier.Copy(&c, req.Comment)

	if err := s.serviceCtx.CommentModel.Insert(ctx, &c); err != nil {
		return nil, err
	}
	var MMDD = "01/02" //MMDD日期格式
	resp = &comment.InsertResp{
		CreateDate: c.CreatedAt.Format(MMDD),
	}
	return resp, nil
}

// Delete implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Delete(ctx context.Context, req *comment.DeleteReq) (resp *comment.DeleteResp, err error) {
	err = s.serviceCtx.CommentModel.Delete(ctx, req.CommentId)
	if err != nil {
		return nil, err
	}
	return &comment.DeleteResp{}, nil
}
