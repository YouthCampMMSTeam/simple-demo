package svcctx

import (
	"douyin-project/microservice/relation/model"
	"douyin-project/microservice/relation/rpc/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	RelationModel model.RelationModel
}

func NewServiceContext(c *config.Config) *ServiceContext {

	//TODO 第二个参数
	conn, _ := gorm.Open(mysql.Open(c.DbSource), &gorm.Config{})
	return &ServiceContext{
		RelationModel: model.NewRelationDbModel(conn),
	}
}
