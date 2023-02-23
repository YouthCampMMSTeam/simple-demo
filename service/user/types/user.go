package types

type Response struct {
	StatusCode int32  `json:status_code`
	StatusMsg  string `json:status_msg`
}
type UserLoginReq struct {
	Name     string `form:"username" json:"username" query:"username"`
	Password string `form:"password" json:"password" query:"password"`
}
type UserLoginResp struct {
	Response
	Token  string
	UserId int64
}

type UserLoginLogicReq struct {
	Name     string
	Password string
}
type UserLoginLogicResp struct {
	UserId int64
}

type UserRegisterReq struct {
	Name     string `form:"username" json:"username" query:"username"`
	Password string `form:"password" json:"password" query:"password"`
}
type UserRegisterResp struct {
	StatusCode int32  `json:status_code`
	StatusMsg  string `json:status_msg`
	Token      string `json:token`
	UserId     int64  `json:user_id`
}
type UserRegisterLogicReq struct {
	Name     string
	Password string
}
type UserRegisterLogicResp struct {
	Token  string
	UserId int64
}

type UserInfoReq struct {
	UserId int64  `form:"user_id" json:"user_id" query:"user_id"`
	Token  string `form:"token" json:"token" query:"token"`
}

type UserInfoResp struct {
	StatusCode int32  `json:status_code`
	StatusMsg  string `json:status_msg`
	User       User
}

type UserInfoLogicReq struct {
	UserId        int64
	CurrentUserId int64
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
