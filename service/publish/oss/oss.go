package oss

import (
	"fmt"
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// TranToUrl file: 原始文件，filename: 新文件名字，例如xxx.mp4,只提供xxx即可
func TranToUrl(file *multipart.FileHeader, filename string) (url string, err error) {
	client, err := oss.New("https://oss-cn-shanghai.aliyuncs.com", "LTAI5t9xU7a8kjDkHNHaQ8ti", "IsCWxAoPUPkpRVa3kihDldswsPFPeO")
	if err != nil {
		return "", err
	}
	// 指定bucket
	bucket, err := client.Bucket("pyp-vedios") // 根据自己的填写
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)

	if err != nil {
		return "", err
	}
	fmt.Println(file.Filename)
	idx := len(file.Filename) - 1
	for ; idx >= 0; idx-- {
		if file.Filename[idx] == '.' {
			break
		}
	}
	path := filename + file.Filename[idx:]
	err = bucket.PutObject(path, src)
	return "https://pyp-vedios.oss-cn-shanghai.aliyuncs.com/" + filename + file.Filename[idx:], err
}
