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
