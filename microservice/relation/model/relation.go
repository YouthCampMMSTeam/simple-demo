package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Relation struct {
	ID         int64 `gorm:"primary_key"`
	FollowId   int64
	FollowerId int64
	IsDeleted  bool
}

//原来默认是relations，这里注意单复数和大小写
func (r *Relation) TableName() string {
	return "Relation"
}

type RelationModel interface {
	FindByFollowidAndFollowerid(ctx context.Context, followId int64, followerId int64) (*Relation, error)
	FindByFollowerid(ctx context.Context, followerId int64) ([]*Relation, error)
	FindByFollowid(ctx context.Context, followId int64) ([]*Relation, error)
}

type relationSqlModel struct {
	SqlConn *gorm.DB
}

func NewRelationDbModel(sqlConn *gorm.DB) RelationModel {
	return &relationSqlModel{
		SqlConn: sqlConn,
	}
}

//查询关注列表（根据两个用户id查询）
func (m *relationSqlModel) FindByFollowidAndFollowerid(ctx context.Context, followId int64, followerId int64) (*Relation, error) {
	var result Relation
	//注意这里where中的命名规则是和数据库一致的，即使用下划线而非大小写
	if err := m.SqlConn.WithContext(ctx).Where("follow_id = ?", followId).Where("follower_id = ?", followerId).First(&result).Error; err != nil {
		// 空结果处理都一样
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

// 查询关注列表（根据关注用户id）
func (m *relationSqlModel) FindByFollowerid(ctx context.Context, followerId int64) ([]*Relation, error) {
	var results []*Relation

	//使用数组查询，空结果也不会有error，而是通过len(results)判断结果个数
	if err := m.SqlConn.WithContext(ctx).Where("follower_id = ?", followerId).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// 查询关注列表（根据被关注用户id）
func (m *relationSqlModel) FindByFollowid(ctx context.Context, followId int64) ([]*Relation, error) {
	var results []*Relation
	if err := m.SqlConn.WithContext(ctx).Where("follow_id = ?", followId).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
