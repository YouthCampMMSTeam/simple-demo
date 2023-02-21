package main

import (
	"context"
	"douyin-project/microservice/user/rpc/kitex_gen/user"
	"douyin-project/microservice/user/rpc/kitex_gen/user/userservice"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	CommentServiceName := "douyin-user-service"
	client, err := userservice.NewClient(
		CommentServiceName,
		//etcd
		client.WithResolver(r),
	)
	if err != nil {
		log.Fatal(err)
	}

	// FindByNameResp FindByName(1: FindByNameRequest req);
	// FindByUserIdResp FindByUserId(1: FindByUserIdRequest req);
	// InsertResp Insert(1: InsertRequest req);
	// UpdateResp Update(1: UpdateRequest req);

	log.Println("FindByName Test Start")

	findReq := &user.FindByNameRequest{
		UserName: "xiaoming",
	}
	findResp, err := client.FindByName(context.Background(), findReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("len(findResp.UserList) %v\n", len(findResp.UserList))
	for ind, comment := range findResp.UserList {
		fmt.Printf("findResp.UserList[%d] %v\n", ind, comment)
	}
	log.Println("FindByName Test End")

	log.Println("FindByUserId Test Start")
	find2Req := &user.FindByUserIdRequest{
		UserId: 1,
	}
	find2Resp, err := client.FindByUserId(context.Background(), find2Req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("len(findResp.UserList) %v\n", len(find2Resp.UserList))
	for ind, comment := range find2Resp.UserList {
		fmt.Printf("findResp.UserList[%d] %v\n", ind, comment)
	}
	log.Println("FindByUserId Test End")

	log.Println("Insert Test Start")
	insertReq := &user.InsertRequest{
		User: &user.User{
			Name:     "老王",
			Password: "123",
		},
	}

	insertResp, err := client.Insert(context.Background(), insertReq)
	if err != nil {
		fmt.Printf("+%v", err)
	} else {
		fmt.Printf("insertResp.UserId %v\n", insertResp.UserId)
	}

	log.Println("Insert Test End")

	log.Println("Update Test Start")

	delReq := &user.UpdateRequest{
		User: &user.User{
			Id:       1,
			Password: "password",
		},
	}

	_, err = client.Update(context.Background(), delReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Update Test End")

}
