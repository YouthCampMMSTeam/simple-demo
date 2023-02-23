namespace go comment

struct Comment {
    1: required i64 Id
    2: required i64 VideoId
    3: required i64 UserId;
    4: required string Content;
    5: required string CreateDate;
}

struct InsertReq {
    1: required Comment comment
}

struct InsertResp {
    1: required string CreateDate
}

struct FindByVideoIdReq {
    1: required i64 VideoId
}

struct FindByVideoIdResp {
    1: list<Comment> commentList
}

struct FindCommentByVideoIdLimit30Req {
    1: required i64 VideoId
}

struct FindCommentByVideoIdLimit30Resp {
    1: list<Comment> commentList
}


struct DeleteReq {
    1: required i64 CommentId
}
struct DeleteResp {
}
service CommentService {
    InsertResp Insert(1: InsertReq req);
    DeleteResp Delete(1: DeleteReq req);
    FindByVideoIdResp FindByVideoId(1: FindByVideoIdReq req);
    FindCommentByVideoIdLimit30Resp FindCommentByVideoIdLimit30(1: FindCommentByVideoIdLimit30Req req);
}

//kitex -module douyin-project -service Comment ./idl/comment.thrift