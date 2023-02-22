package svcctx

import (
	"douyin-project/microservice/user/rpc/kitex_gen/user/userservice"
	"douyin-project/service/user/config"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type ServiceContext struct {
	// EtcdResolver etcd.discovery.Resolver
	UserRpc userservice.Client
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

	return &ServiceContext{
		// EtcdResolver: ,
		UserRpc: userRpcClient,
	}
}
