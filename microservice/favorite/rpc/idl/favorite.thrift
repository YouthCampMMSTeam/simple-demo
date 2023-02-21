namespace go favorite

struct Favorite {
    1: required i64 Id
    2: required i64 VideoId
    3: required i64 UserId;
}

struct FindByVideoIdAndUserIdRequest {
    1: required i64 VideoId
    2: required i64 UserId;
}

struct FindByVideoIdAndUserIdResp {
    1: list<Favorite> favoriteList
}

struct FindByUserIdRequest {
    1: required i64 UserId
}

struct FindByUserIdResp {
    1: list<Favorite> favoriteList
}

struct InsertRequest {
    1: required Favorite favorite
}

struct InsertResp {
    1: required i64 favoriteId
}

struct DeleteRequest {
    1: required i64 favoriteId
}

struct DeleteResp {
}


service FavoriteService {
    FindByVideoIdAndUserIdResp FindByVideoIdAndUserId(1: FindByVideoIdAndUserIdRequest req);
    FindByUserIdResp FindByUserId(1: FindByUserIdRequest req);
    InsertResp Insert(1: InsertRequest req);
    DeleteResp Delete(1: DeleteRequest req);
}

//kitex -module douyin-project -service Comment ./idl/favorite.thrift