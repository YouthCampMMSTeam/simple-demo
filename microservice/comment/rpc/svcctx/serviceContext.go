package svcctx

import (
	"douyin-project/microservice/comment/model"
	"douyin-project/microservice/comment/rpc/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	CommentModel model.CommentModel
}

func NewServiceContext(c *config.Config) *ServiceContext {
	//TODO 第二个参数
	conn, _ := gorm.Open(mysql.Open(c.DbSource), &gorm.Config{})
	return &ServiceContext{
		CommentModel: model.NewCommentSqlModel(conn),
	}
}
