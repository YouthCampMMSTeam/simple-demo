package message

import "douyin-project/microservice/video/rpc/kitex_gen/video"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type VideoListResponse struct {
	Response
	VideoList []*video.Video `json:"video_list"`
}
