package svcctx

import (
	"douyin-project/microservice/favorite/model"
	"douyin-project/microservice/favorite/rpc/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	FavoriteModel model.FavoriteModel
}

func NewServiceContext(c *config.Config) *ServiceContext {
	//TODO 第二个参数
	conn, _ := gorm.Open(mysql.Open(c.DbSource), &gorm.Config{})
	return &ServiceContext{
		FavoriteModel: model.NewFavoriteSqlModel(conn),
	}
}
