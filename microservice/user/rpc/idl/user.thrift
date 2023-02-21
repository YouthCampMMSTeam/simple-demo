namespace go user

struct User {
    1: required i64 Id
    2: required string Name
    3: required string Password;
    4: required i64 FollowCount;
    5: required i64 FollowerCount;
}

struct FindByNameRequest {
    1: required string userName
}

struct FindByNameResp {
    1: list<User> userList
}
struct FindByUserIdRequest {
    1: required i64 userId
}

struct FindByUserIdResp {
    1: list<User> userList
}

struct InsertRequest {
    1: required User User
}
struct InsertResp {
    1: required i64 userId
}
struct UpdateRequest {
    1: required User User
}
struct UpdateResp {
}

service UserService {
    FindByNameResp FindByName(1: FindByNameRequest req);
    FindByUserIdResp FindByUserId(1: FindByUserIdRequest req);
    InsertResp Insert(1: InsertRequest req);
    UpdateResp Update(1: UpdateRequest req);
}

//kitex -module douyin-project -service User ./idl/user.thrift