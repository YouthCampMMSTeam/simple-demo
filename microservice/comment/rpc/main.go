package main

import (
	"douyin-project/microservice/comment/rpc/config"
	"log"
	"net"

	comment "douyin-project/microservice/comment/rpc/kitex_gen/comment/commentservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	c, err := config.Parse("./config/comment.yaml")
	if err != nil {
		panic(err)
	}

	//etcd
	r, err := etcd.NewEtcdRegistry([]string{c.EtcdAdress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", c.ServiceAddress)
	serviceImp := NewCommentServiceImpl(c)
	print(c.ServiceName)
	svr := comment.NewServer(serviceImp,
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
