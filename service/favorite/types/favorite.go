package types

type Response struct {
	StatusCode int32  `json:status_code`
	StatusMsg  string `json:status_msg`
}

type FavoriteActionReq struct {
	// Token      string `form:"token" json:"token" query:"token"`
	VideoId    int64 `form:"video_id" json:"video_id" query:"video_id"`
	ActionType int32 `form:"action_type" json:"action_type" query:"action_type"`
}
type FavoriteActionResp struct {
	StatusCode int32  `json:status_code`
	StatusMsg  string `json:status_msg`
}

type FavoriteActionLogicReq struct {
	UserId     int64
	VideoId    int64
	ActionType int32
}
type FavoriteActionLogicResp struct {
}

type FavoriteListReq struct {
	Token  string `form:"token" json:"token" query:"token"`
	UserId int64  `form:"user_id" json:"user_id" query:"user_id"`
}
type FavoriteListLogicReq struct {
	CurrentUserId int64
	UserId        int64
}
type FavoriteListResp struct {
	StatusCode int32    `json:status_code`
	StatusMsg  string   `json:status_msg`
	VideoList  []*Video `json:"video_list"`
}
type Video struct {
	Id            int64  `json:"id"`
	Author        *User  `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
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
	// TotalFavorited  int64  `json:"total_favorited"`
	// WorkCount       int64  `json:"work_count"`
	// FavoriteCount   int64  `json:"favorite_count"`
}
