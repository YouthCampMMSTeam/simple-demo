package types

type Response struct {
	StatusCode int32  `json:status_code`
	StatusMsg  string `json:status_msg`
}

type FeedReq struct {
	Token      string `form:"token" json:"token" query:"token"`
	LatestTime int64  `form:"latest_time" json:"latest_time" query:"latest_time"`
}
type FeedResp struct {
	StatusCode int32    `json:status_code`
	StatusMsg  string   `json:status_msg`
	VideoList  []*Video `json:video_list`
	NextTime   int64    `json:next_time`
}

type FeedLogicReq struct {
	CurrentUserId int64
	LatestTime    int64
}
type Video struct {
	Id            int64  `json:"id"`
	Author        *User  `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentConut  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	// Avatar          string `json:"avatar"`
	// BackgroundImage string `json:"background_image"`
	// Signature       string `json:"signature"`
	// TotalPublishd  int64  `json:"total_Publishd"`
	// WorkCount       int64  `json:"work_count"`
	// PublishCount   int64  `json:"Publish_count"`
}
