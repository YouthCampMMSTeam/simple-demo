package types

type Response struct {
	StatusCode int32  `json:status_code`
	StatusMsg  string `json:status_msg`
}

type CommentActionReq struct {
	Token       string `form:"token" json:"token" query:"token"`
	VideoId     int64  `form:"video_id" json:"video_id" query:"video_id"`
	ActionType  int32  `form:"action_type" json:"action_type" query:"action_type"`
	CommentText string `form:"comment_text" json:"comment_text" query:"comment_text"`
	CommentId   int64  `form:"comment_id" json:"comment_id" query:"comment_id"`
}
type CommentActionResp struct {
	StatusCode int32   `json:status_code`
	StatusMsg  string  `json:status_msg`
	Comment    Comment `json:comment`
}

type CommentActionLogicReq struct {
	CurrentUserId int64
	VideoId       int64
	ActionType    int32
	CommentText   string
	CommentId     int64
}

type CommentActionLogicResp struct {
	Comment Comment
}

type CommentListReq struct {
	Token   string `form:"token" json:"token" query:"token"`
	VideoId int64  `form:"video_id" json:"video_id" query:"video_id"`
}
type CommentListResp struct {
	StatusCode  int32      `json:status_code`
	StatusMsg   string     `json:status_msg`
	CommentList []*Comment `json:"comment_list"`
}

type CommentListLogicReq struct {
	CurrentUserId int64
	VideoId       int64
}

// type CommentListLogicResp struct {
// 	CommentList []*Comment
// }
type Comment struct {
	Id         int64
	User       *User
	Content    string
	CreateDate string
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
	// TotalCommentd  int64  `json:"total_commentd"`
	// WorkCount       int64  `json:"work_count"`
	// CommentCount   int64  `json:"comment_count"`
}
