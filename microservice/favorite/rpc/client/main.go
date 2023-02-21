package main

import (
	"context"
	"douyin-project/microservice/favorite/rpc/kitex_gen/favorite"
	"douyin-project/microservice/favorite/rpc/kitex_gen/favorite/favoriteservice"
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

	CommentServiceName := "douyin-favorite-service"
	client, err := favoriteservice.NewClient(
		CommentServiceName,
		//etcd
		client.WithResolver(r),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("FindByVideoIdAndUserId Test Start")

	findReq := &favorite.FindByVideoIdAndUserIdRequest{
		VideoId: 1,
		UserId:  1,
	}
	findResp, err := client.FindByVideoIdAndUserId(context.Background(), findReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("len(findResp.FavoriteList) %v\n", len(findResp.FavoriteList))
	for ind, comment := range findResp.FavoriteList {
		fmt.Printf("findResp.FavoriteList[%d] %v\n", ind, comment)
	}
	log.Println("FindByVideoIdAndUserId Test End")

	log.Println("FindByUserId Test Start")

	find2Req := &favorite.FindByUserIdRequest{
		UserId: 2,
	}
	find2Resp, err := client.FindByUserId(context.Background(), find2Req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("len(findResp.FavoriteList) %v\n", len(find2Resp.FavoriteList))
	for ind, comment := range find2Resp.FavoriteList {
		fmt.Printf("findResp.FavoriteList[%d] %v\n", ind, comment)
	}
	log.Println("FindByUserId Test End")

	// log.Println("Insert Test Start")

	// insertReq := &favorite.InsertRequest{
	// 	Favorite: &favorite.Favorite{
	// 		VideoId: 2,
	// 		UserId:  2,
	// 	},
	// }

	// insertResp, err := client.Insert(context.Background(), insertReq)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("insertResp.FavoriteId %v\n", insertResp.FavoriteId)
	// log.Println("Insert Test End")

	// log.Println("Delete Test Start")

	// delReq := &favorite.DeleteRequest{
	// 	FavoriteId: 1,
	// }

	// _, err = client.Delete(context.Background(), delReq)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Delete Test End")

}
