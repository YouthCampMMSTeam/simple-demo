package model

import (
	"context"

	"gorm.io/gorm"
)

//这里定义顺序要和数据库中字段定义顺序一致
type User struct {
	// gorm.Model
	ID int64 `gorm:"primary_key" copier:"Id"` //kitex中是Id，所以得设置了copier才能用
	//考虑到gorm.Model中间是ID，所以这里不能设置
	Name          string
	Password      string
	FollowCount   uint64
	FollowerCount uint64
}

//原来默认是relations，这里注意单复数和大小写
func (u *User) TableName() string {
	return "User"
}

type UserModel interface {
	FindByName(ctx context.Context, Name string) ([]*User, error)
	FindByUserId(ctx context.Context, UserId int64) ([]*User, error)
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}

type userSqlModel struct {
	SqlConn *gorm.DB
}

func NewUserSqlModel(sqlConn *gorm.DB) UserModel {
	return &userSqlModel{
		SqlConn: sqlConn,
	}
}

func (m *userSqlModel) FindByName(ctx context.Context, Name string) ([]*User, error) {
	var results []*User
	//注意这里where中的命名规则是和数据库一致的，即使用下划线而非大小写
	if err := m.SqlConn.WithContext(ctx).Where("name = ?", Name).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil

}
func (m *userSqlModel) FindByUserId(ctx context.Context, UserId int64) ([]*User, error) {
	var results []*User
	//注意这里where中的命名规则是和数据库一致的，即使用下划线而非大小写
	if err := m.SqlConn.WithContext(ctx).Where("id = ?", UserId).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
func (m *userSqlModel) Insert(ctx context.Context, user *User) error {
	if err := m.SqlConn.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *userSqlModel) Update(ctx context.Context, user *User) error {
	userMod := &User{
		ID: user.ID,
	}
	if err := m.SqlConn.WithContext(ctx).Model(userMod).Updates(*user).Error; err != nil {
		return err
	}
	return nil
}
