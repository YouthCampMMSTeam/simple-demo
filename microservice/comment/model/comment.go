package model

import (
	"context"

	"gorm.io/gorm"
)

//这里定义顺序要和数据库中字段定义顺序一致
type Comment struct {
	gorm.Model
	// Id      int64  `gorm:"primary_key"`
	VideoId uint64 `json:"video_id"`
	UserId  uint64 `json:"user_id"`
	Content string `json:"content"`
}

//原来默认是relations，这里注意单复数和大小写
func (c *Comment) TableName() string {
	return "Comment"
}

type CommentModel interface {
	FindByVideoIdLimit30(ctx context.Context, videoId int64) ([]*Comment, error)
	FindByVideoId(ctx context.Context, videoId int64) ([]*Comment, error)
	Insert(ctx context.Context, comment *Comment) error
	Delete(ctx context.Context, commentId int64) error
}

type commentSqlModel struct {
	SqlConn *gorm.DB
}

func NewCommentSqlModel(sqlConn *gorm.DB) CommentModel {
	return &commentSqlModel{
		SqlConn: sqlConn,
	}
}

//CommentIdList --> FindByVideoId
//根据视频id获取评论id 列表
func (m *commentSqlModel) FindByVideoIdLimit30(ctx context.Context, videoId int64) ([]*Comment, error) {
	var results []*Comment
	//注意这里where中的命名规则是和数据库一致的，即使用下划线而非大小写
	if err := m.SqlConn.WithContext(ctx).Where("video_id = ?", videoId).Order("created_at desc").Limit(30).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (m *commentSqlModel) FindByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	var results []*Comment
	//注意这里where中的命名规则是和数据库一致的，即使用下划线而非大小写
	if err := m.SqlConn.WithContext(ctx).Where("video_id = ?", videoId).Order("created_at desc").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

//增加评论
func (m *commentSqlModel) Insert(ctx context.Context, comment *Comment) error {
	if err := m.SqlConn.WithContext(ctx).Create(comment).Error; err != nil {
		return err
	}
	return nil
}

// 删除评论
func (m *commentSqlModel) Delete(ctx context.Context, commentId int64) error {
	//使用数组查询，空结果也不会有error，而是通过len(results)判断结果个数
	if err := m.SqlConn.WithContext(ctx).Delete(&Comment{}, commentId).Error; err != nil {
		return err
	}
	return nil
}
