package main

import (
	"context"

	"douyin-project/microservice/relation/rpc/config"
	"douyin-project/microservice/relation/rpc/impl"
	"douyin-project/microservice/relation/rpc/kitex_gen/relation"
	"douyin-project/microservice/relation/rpc/svcctx"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct {
	serviceCtx *svcctx.ServiceContext
}

func NewRelationServiceImpl(c *config.Config) *RelationServiceImpl {
	return &RelationServiceImpl{
		serviceCtx: svcctx.NewServiceContext(c),
	}
}

// SelectRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) SelectRelation(ctx context.Context, req *relation.SelectRelationRequest) (resp *relation.SelectRelationResp, err error) {
	// TODO: Your code here...
	i := impl.NewRelationImpl(ctx, s.serviceCtx)
	return i.SelectRelation(req)
}
