package main

import (
	"douyin-project/microservice/favorite/rpc/config"
	"log"
	"net"

	favorite "douyin-project/microservice/favorite/rpc/kitex_gen/favorite/favoriteservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	c, err := config.Parse("./config/favorite.yaml")
	if err != nil {
		panic(err)
	}

	//etcd
	r, err := etcd.NewEtcdRegistry([]string{c.EtcdAdress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", c.ServiceAddress)
	serviceImp := NewFavoriteServiceImpl(c)
	print(c.ServiceName)
	svr := favorite.NewServer(serviceImp,
		//server地址
		server.WithServiceAddr(addr),
		//etcd
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: c.ServiceName}),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
