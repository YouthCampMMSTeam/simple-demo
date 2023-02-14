package impl

import (
	"context"

	"douyin-project/microservice/relation/model"
	"douyin-project/microservice/relation/rpc/kitex_gen/relation"
	"douyin-project/microservice/relation/rpc/svcctx"

	"github.com/jinzhu/copier"
)

type RelationImpl struct {
	ctx        context.Context
	serviceCtx *svcctx.ServiceContext
}

func NewRelationImpl(ctx context.Context, serviceCtx *svcctx.ServiceContext) *RelationImpl {
	return &RelationImpl{
		ctx:        ctx,
		serviceCtx: serviceCtx,
	}
}

func (i *RelationImpl) SelectRelation(req *relation.SelectRelationRequest) (*relation.SelectRelationResp, error) {
	// var results []model.Relation
	results := []*model.Relation{}
	if req.FollowerId == -1 && req.FollowId == -1 {
		//相当于两个参数都没有传入，自然没有结果
		results = []*model.Relation{}
	} else if req.FollowerId == -1 {
		ret, err := i.serviceCtx.RelationModel.SelectRelationByFollowid(i.ctx, req.FollowId)
		if err != nil {
			return nil, err
		}
		results = ret
	} else if req.FollowId == -1 {
		ret, err := i.serviceCtx.RelationModel.SelectRelationByFollowerid(i.ctx, req.FollowerId)
		if err != nil {
			return nil, err
		}
		results = ret
	} else {
		result, err := i.serviceCtx.RelationModel.SelectRelationByFollowidAndFollowerid(i.ctx, req.FollowId, req.FollowerId)
		if err != nil {
			return nil, err
		}
		if result == nil {
			results = []*model.Relation{} //没有则包装空数组，后续通过数组长度判断是否为空
		} else {
			results = []*model.Relation{result}
		}
	}
	resp := &relation.SelectRelationResp{}
	//copier针对切片要使用&
	copier.Copy(&resp.RelationList, &results)

	return resp, nil
}
