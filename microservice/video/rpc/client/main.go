package main

import (
	"context"
	"douyin-project/microservice/video/rpc/kitex_gen/video"
	"douyin-project/microservice/video/rpc/kitex_gen/video/videoservice"
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

	CommentServiceName := "douyin-video-service"
	client, err := videoservice.NewClient(
		CommentServiceName,
		//etcd
		client.WithResolver(r),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("FindOrderByTime Test Start")
	findResp, err := client.FindOrderByTime(context.Background(), &video.FindOrderByTimeReq{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("len(findResp.VideoList) %v\n", len(findResp.VideoList))
	for ind, comment := range findResp.VideoList {
		fmt.Printf("findResp.VideoList[%d] %v\n", ind, comment)
	}
	log.Println("FindOrderByTime Test End")

	log.Println("FindByVideoId Test Start")
	find2Req := &video.FindByVideoIdReq{
		VideoId: 1,
	}
	find2Resp, err := client.FindByVideoId(context.Background(), find2Req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("len(findResp.VideoList) %v\n", len(find2Resp.VideoList))
	for ind, comment := range find2Resp.VideoList {
		fmt.Printf("findResp.VideoList[%d] %v\n", ind, comment)
	}
	log.Println("FindByVideoId Test End")

	log.Println("FindByUserId Test Start")
	find3Req := &video.FindByUserIdReq{
		UserId: 1,
	}
	find3Resp, err := client.FindByUserId(context.Background(), find3Req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("len(findResp.VideoList) %v\n", len(find3Resp.VideoList))
	for ind, comment := range find3Resp.VideoList {
		fmt.Printf("findResp.VideoList[%d] %v\n", ind, comment)
	}
	log.Println("FindByUserId Test End")

	log.Println("Insert Test Start")
	insertReq := &video.InsertReq{
		Video: &video.Video{
			PlayUrl:       "dafdsafad",
			CoverUrl:      "dafdsa",
			FavoriteCount: 0,
			CommentCount:  0,
			AuthorId:      1,
		},
	}

	insertResp, err := client.Insert(context.Background(), insertReq)
	if err != nil {
		fmt.Printf("+%v", err)
	} else {
		fmt.Printf("insertResp.VideoId %v\n", insertResp.VideoId)
	}
	log.Println("Insert Test End")

	log.Println("Update Test Start")
	delReq := &video.UpdateReq{
		Video: &video.Video{
			Id:            2,
			PlayUrl:       "aaaaaaaaaa",
			CoverUrl:      "sdafasd",
			FavoriteCount: 0,
			CommentCount:  0,
			AuthorId:      1,
		},
	}
	_, err = client.Update(context.Background(), delReq)
	if err != nil {
		log.Println(err)
	}
	log.Println("Update Test End")
}
