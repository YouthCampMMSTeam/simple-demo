namespace go relation

struct Relation {
    1: required i64 Id
    2: required i64 FollowId
    3: required i64 FollowerId;
    4: required bool IsDeleted;
}

struct SelectRelationRequest {
    1: required i64 followId
    2: required i64 followerId
}

struct SelectRelationResp {
    1: list<Relation> relationList
}

service RelationService {
    SelectRelationResp SelectRelation(1: SelectRelationRequest req);
}

//kitex -module douyin-project -service Relation relation.thrift