package model

import (
	"context"

	"gorm.io/gorm"
)

//这里定义顺序要和数据库中字段定义顺序一致
type Favorite struct {
	gorm.Model
	VideoId uint64 `json:"video_id"`
	UserId  uint64 `json:"user_id"`
}

//原来默认是relations，这里注意单复数和大小写
func (f *Favorite) TableName() string {
	return "Favorite"
}

type FavoriteModel interface {
	FindByVideoIdAndUserId(ctx context.Context, videoId int64, userId int64) ([]*Favorite, error)
	FindByUserId(ctx context.Context, userId int64) ([]*Favorite, error)
	Insert(ctx context.Context, favorite *Favorite) error
	// Update(ctx context.Context, favorite *Favorite) error
	Delete(ctx context.Context, favoriteId int64) error
}

type favoriteSqlModel struct {
	SqlConn *gorm.DB
}

func NewFavoriteSqlModel(sqlConn *gorm.DB) FavoriteModel {
	return &favoriteSqlModel{
		SqlConn: sqlConn,
	}
}

func (m *favoriteSqlModel) FindByVideoIdAndUserId(ctx context.Context, videoId int64, userId int64) ([]*Favorite, error) {
	var results []*Favorite
	if err := m.SqlConn.WithContext(ctx).Where("video_id = ?", videoId).Where("user_id = ?", userId).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
func (m *favoriteSqlModel) FindByUserId(ctx context.Context, userId int64) ([]*Favorite, error) {
	var results []*Favorite
	if err := m.SqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
func (m *favoriteSqlModel) Insert(ctx context.Context, favorite *Favorite) error {
	if err := m.SqlConn.WithContext(ctx).Create(favorite).Error; err != nil {
		return err
	}
	return nil
}

// func (m *favoriteSqlModel) Update(ctx context.Context, favorite *Favorite) error {
// 	if err := m.SqlConn.WithContext(ctx).Create(favorite).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func (m *favoriteSqlModel) Delete(ctx context.Context, favoriteId int64) error {
	if err := m.SqlConn.WithContext(ctx).Delete(&Favorite{}, favoriteId).Error; err != nil {
		return err
	}
	return nil
}
