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

service VideoService {
    FindOrderByTimeResp FindOrderByTime(1: FindOrderByTimeReq req);
    // FindOrderByTimeRangeResp FindOrderByTimeRange(1: FindOrderByTimeRangeReq req);
    FindByVideoIdResp FindByVideoId(1: FindByVideoIdReq req);
    FindByUserIdResp FindByUserId(1: FindByUserIdReq req);
    InsertResp Insert(1: InsertReq req);
    UpdateResp Update(1: UpdateReq req);
}

//kitex -module douyin-project -service Video ./idl/video.thrift
