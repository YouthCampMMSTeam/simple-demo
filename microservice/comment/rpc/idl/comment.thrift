namespace go comment

struct Comment {
    1: required i64 Id
    2: required i64 VideoId
    3: required i64 UserId;
    4: required string Content;
}

struct InsertRequest {
    1: required Comment comment
}

struct InsertResp {
    1: required i64 CommentId
}

struct FindCommentByVideoIdLimit30Request {
    1: required i64 VideoId
}

struct FindCommentByVideoIdLimit30Resp {
    1: list<Comment> commentList
}

service CommentService {
    InsertResp Insert(1: InsertRequest req);
    FindCommentByVideoIdLimit30Resp FindCommentByVideoIdLimit30(1: FindCommentByVideoIdLimit30Request req);
}

//kitex -module douyin-project -service Comment ./idl/comment.thrift