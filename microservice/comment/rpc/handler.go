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
func (s *CommentServiceImpl) FindCommentByVideoIdLimit30(ctx context.Context, req *comment.FindCommentByVideoIdLimit30Request) (resp *comment.FindCommentByVideoIdLimit30Resp, err error) {
	results, err := s.serviceCtx.CommentModel.FindByVideoIdLimit30(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	resp = &comment.FindCommentByVideoIdLimit30Resp{}
	copier.Copy(&resp.CommentList, &results)
	return resp, nil
}

// Insert implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Insert(ctx context.Context, req *comment.InsertRequest) (resp *comment.InsertResp, err error) {
	var c model.Comment
	copier.Copy(&c, req.Comment)

	if err := s.serviceCtx.CommentModel.Insert(ctx, &c); err != nil {
		return nil, err
	}
	resp = &comment.InsertResp{
		CommentId: int64(c.ID),
	}
	return resp, nil
}
