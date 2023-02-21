package main

import (
	"log"
	"net"

	"douyin-project/microservice/user/rpc/config"
	user "douyin-project/microservice/user/rpc/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	c, err := config.Parse("./config/User.yaml")
	if err != nil {
		panic(err)
	}

	//etcd
	r, err := etcd.NewEtcdRegistry([]string{c.EtcdAdress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", c.ServiceAddress)
	serviceImp := NewUserServiceImpl(c)
	print(c.ServiceName)
	svr := user.NewServer(serviceImp,
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
