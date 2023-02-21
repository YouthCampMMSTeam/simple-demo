package svcctx

import (
	"douyin-project/microservice/user/model"
	"douyin-project/microservice/user/rpc/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	UserModel model.UserModel
}

func NewServiceContext(c *config.Config) *ServiceContext {
	//TODO 第二个参数
	conn, _ := gorm.Open(mysql.Open(c.DbSource), &gorm.Config{})
	return &ServiceContext{
		UserModel: model.NewUserSqlModel(conn),
	}
}
