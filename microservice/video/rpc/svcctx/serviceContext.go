package svcctx

import (
	"douyin-project/microservice/video/model"
	"douyin-project/microservice/video/rpc/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	VideoModel model.VideoModel
}

func NewServiceContext(c *config.Config) *ServiceContext {
	//TODO 第二个参数
	conn, _ := gorm.Open(mysql.Open(c.DbSource), &gorm.Config{})
	return &ServiceContext{
		VideoModel: model.NewVideoSqlModel(conn),
	}
}
