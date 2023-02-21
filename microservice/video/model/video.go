package model

import (
	"context"

	"gorm.io/gorm"
)

//这里定义顺序要和数据库中字段定义顺序一致
type Video struct {
	gorm.Model
	PlayUrl       string
	CoverUrl      string
	FavoriteCount uint64
	CommentCount  uint64
	Title         string
	AuthorId      uint64
}

//原来默认是relations，这里注意单复数和大小写
func (v *Video) TableName() string {
	return "Video"
}

type VideoModel interface {
	FindOrderByTime(ctx context.Context, limitNum int64) ([]*Video, error)
	// FindOrderByTimeRange(ctx context.Context, earliestTime time.Time, latestTime time.Time) ([]*Video, error)
	FindByVideoId(ctx context.Context, videoId int64) ([]*Video, error)
	FindByUserId(ctx context.Context, userId int64) ([]*Video, error)
	Insert(ctx context.Context, video *Video) error
	Update(ctx context.Context, video *Video) error
}

type videoSqlModel struct {
	SqlConn *gorm.DB
}

func NewVideoSqlModel(sqlConn *gorm.DB) VideoModel {
	return &videoSqlModel{
		SqlConn: sqlConn,
	}
}

func (m *videoSqlModel) FindOrderByTime(ctx context.Context, limitNum int64) ([]*Video, error) {
	var results []*Video
	if err := m.SqlConn.WithContext(ctx).Order("created_at desc").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// func (m *videoSqlModel) FindOrderByTimeRange(ctx context.Context, earliestTime time.Time, latestTime time.Time) ([]*Video, error) {
// 	var results []*Video
// 	if err := m.SqlConn.WithContext(ctx).Where("created_at >= ? AND created_at <= ?", earliestTime, latestTime).Order("created_at desc").Find(&results).Error; err != nil {
// 		return nil, err
// 	}
// 	return results, nil
// }
func (m *videoSqlModel) FindByVideoId(ctx context.Context, videoId int64) ([]*Video, error) {
	var results []*Video
	if err := m.SqlConn.WithContext(ctx).Where("id = ?", videoId).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
func (m *videoSqlModel) FindByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	var results []*Video
	if err := m.SqlConn.WithContext(ctx).Where("author_id = ?", userId).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (m *videoSqlModel) Insert(ctx context.Context, video *Video) error {
	if err := m.SqlConn.WithContext(ctx).Create(video).Error; err != nil {
		return err
	}
	return nil
}
func (m *videoSqlModel) Update(ctx context.Context, video *Video) error {
	videoMod := &Video{}
	videoMod.ID = video.ID
	if err := m.SqlConn.WithContext(ctx).Model(videoMod).Updates(*video).Error; err != nil {
		return err
	}
	return nil
}