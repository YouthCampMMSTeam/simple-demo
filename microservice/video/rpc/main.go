package main

import (
	"log"
	"net"

	"douyin-project/microservice/video/rpc/config"
	video "douyin-project/microservice/video/rpc/kitex_gen/video/videoservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	c, err := config.Parse("./config/Video.yaml")
	if err != nil {
		panic(err)
	}

	//etcd
	r, err := etcd.NewEtcdRegistry([]string{c.EtcdAdress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", c.ServiceAddress)
	serviceImp := NewVideoServiceImpl(c)
	print(c.ServiceName)
	svr := video.NewServer(serviceImp,
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
