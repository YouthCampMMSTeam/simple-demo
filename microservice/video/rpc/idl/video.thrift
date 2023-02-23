namespace go video

struct Video {
    1: required i64 Id
    2: required string PlayUrl
    3: required string CoverUrl
    4: required i64 FavoriteCount
    5: required i64 CommentCount
    6: required i64 AuthorId;
}

struct FindOrderByTimeReq {
    1: required i64 limitNum
}

struct FindOrderByTimeResp {
    1: list<Video> videoList
}

struct FindWithTimeLimitReq {
    1: required i64 LatestTime
}

struct FindWithTimeLimitResp {
    1: list<Video> videoList
    2: required i64 NextTime
}


struct FindByVideoIdReq {
    1: required i64 VideoId
}
struct FindByVideoIdResp {
    1: list<Video> videoList
}

struct FindByUserIdReq {
    1: required i64 UserId
}
struct FindByUserIdResp {
    1: list<Video> videoList
}

struct InsertReq {
    1: required Video video
}

struct InsertResp {
    1: required i64 videoId
}
struct UpdateReq {
    1: required Video video
}

struct UpdateResp {
}

struct FavoriteCountModifiedReq {
    1: required i64 videoId
    2: required bool posOrNeg
}

struct FavoriteCountModifiedResp {
}

struct CommentCountModifiedReq {
    1: required i64 videoId
    2: required bool posOrNeg
}

struct CommentCountModifiedResp {
}
service VideoService {
    FindOrderByTimeResp FindOrderByTime(1: FindOrderByTimeReq req);
    // FindOrderByTimeRangeResp FindOrderByTimeRange(1: FindOrderByTimeRangeReq req);
    FindWithTimeLimitResp FindWithTimeLimit(1: FindWithTimeLimitReq req);
    FindByVideoIdResp FindByVideoId(1: FindByVideoIdReq req);
    FindByUserIdResp FindByUserId(1: FindByUserIdReq req);
    InsertResp Insert(1: InsertReq req);
    UpdateResp Update(1: UpdateReq req);
    FavoriteCountModifiedResp FavoriteCountModified(1: FavoriteCountModifiedReq req);
    CommentCountModifiedResp CommentCountModified(1: CommentCountModifiedReq req);
}

//kitex -module douyin-project -service Video ./idl/video.thrift

