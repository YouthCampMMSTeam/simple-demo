package svcctx

import (
	"douyin-project/microservice/favorite/rpc/kitex_gen/favorite/favoriteservice"
	"douyin-project/microservice/relation/rpc/kitex_gen/relation/relationservice"
	"douyin-project/microservice/user/rpc/kitex_gen/user/userservice"
	"douyin-project/microservice/video/rpc/kitex_gen/video/videoservice"
	"douyin-project/service/publish/config"
	"douyin-project/service/publish/uuid"
	"log"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type ServiceContext struct {
	UserRpc userservice.Client
	// VideoRpc    videoservice.Client
	// CommentRpc  commentservice.Client
	RelationRpc relationservice.Client
	FavoriteRpc favoriteservice.Client
	VideoRpc    videoservice.Client
	IdentityKey string

	UuidGenerator *snowflake.Node
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

	favoriteRpcClient, err := favoriteservice.NewClient(
		c.FavoriteServiceName,
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
	// commentRpcClient, err := commentservice.NewClient(
	// 	c.CommentServiceName,
	// 	client.WithResolver(r),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// videoRpcClient, err := videoservice.NewClient(
	// 	c.VideoServiceName,
	// 	client.WithResolver(r),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	uuidGenerator, _ := uuid.NewSnowflake("2021-12-03", 1)
	return &ServiceContext{
		// EtcdResolver: ,
		FavoriteRpc: favoriteRpcClient,
		UserRpc:     userRpcClient,
		// VideoRpc:    videoRpcClient,
		RelationRpc: relationRpcClient,
		// CommentRpc:  commentRpcClient,
		VideoRpc:      videoRpcClient,
		IdentityKey:   c.IdentityKey,
		UuidGenerator: uuidGenerator,
	}
}
