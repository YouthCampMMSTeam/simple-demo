package svcctx

import (
	"douyin-project/microservice/comment/rpc/kitex_gen/comment/commentservice"
	"douyin-project/microservice/relation/rpc/kitex_gen/relation/relationservice"
	"douyin-project/microservice/user/rpc/kitex_gen/user/userservice"
	"douyin-project/microservice/video/rpc/kitex_gen/video/videoservice"
	"douyin-project/service/comment/config"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type ServiceContext struct {
	UserRpc     userservice.Client
	VideoRpc    videoservice.Client
	CommentRpc  commentservice.Client
	RelationRpc relationservice.Client
	IdentityKey string
	//rpc连接
}

func NewServiceContext(c *config.Config) *ServiceContext {
	r, err := etcd.NewEtcdResolver([]string{c.EtcdAdress})
	if err != nil {
		log.Fatal(err)
	}

	userRpcClient, err := userservice.NewClient(
		c.UserServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		log.Fatal(err)
	}

	videoRpcClient, err := videoservice.NewClient(
		c.VideoServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		log.Fatal(err)
	}

	relationRpcClient, err := relationservice.NewClient(
		c.RelationServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	commentRpcClient, err := commentservice.NewClient(
		c.CommentServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceContext{
		// EtcdResolver: ,
		// FavoriteRpc: favoriteRpcClient,
		UserRpc:     userRpcClient,
		VideoRpc:    videoRpcClient,
		RelationRpc: relationRpcClient,
		CommentRpc:  commentRpcClient,
		IdentityKey: c.IdentityKey,
	}
}
