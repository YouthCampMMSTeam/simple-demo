package main

import (
	"context"
	"douyin-project/microservice/comment/rpc/kitex_gen/comment"
	"douyin-project/microservice/comment/rpc/kitex_gen/comment/commentservice"
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

	CommentServiceName := "douyin-comment-service"
	client, err := commentservice.NewClient(
		CommentServiceName,
		//etcd
		client.WithResolver(r),
	)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("InsertCommentRequest Test")

	// insertReq := &comment.InsertCommentReq{
	// 	Comment: &comment.Comment{
	// 		VideoId: 1,
	// 		UserId:  1,
	// 		Content: "我是阳光快乐大男孩",
	// 	},
	// }

	// insertResp, err := client.InsertComment(context.Background(), insertReq)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("resp.CommentId: %v\n", insertResp.CommentId)

	fmt.Println("FindCommentByVideoIdLimit30Request Test")

	findReq := &comment.FindCommentByVideoIdLimit30Req{
		VideoId: 1,
	}

	findResp, err := client.FindCommentByVideoIdLimit30(context.Background(), findReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("len(findResp.CommentList) %v\n", len(findResp.CommentList))
	for ind, comment := range findResp.CommentList {
		fmt.Printf("findResp.CommentList[%d] %v\n", ind, comment)
	}

}
