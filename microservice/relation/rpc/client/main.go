package main

import (
	"context"
	"fmt"
	"log"

	"douyin-project/microservice/relation/rpc/kitex_gen/relation"
	"douyin-project/microservice/relation/rpc/kitex_gen/relation/relationservice"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	RelationServiceName := "douyin-relation-service"
	client, err := relationservice.NewClient(
		RelationServiceName,
		//etcd
		client.WithResolver(r),
	)
	if err != nil {
		log.Fatal(err)
	}

	req := &relation.SelectRelationRequest{
		FollowId:   2,
		FollowerId: -1,
	}

	resp, err := client.SelectRelation(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(resp.RelationList))
	fmt.Println(resp.RelationList)
}
