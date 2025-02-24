mongosh -- "$MONGO_INITDB_DATABASE" <<EOF
db = db.getSiblingDB('$MONGO_INITDB_DATABASE')
db.auth('$MONGO_USERNAME', '$MONGO_PASSWORD');


# 用户表为用户名创建唯一升序索引
db.user.createIndex({"username":1},{"unique":true});

let AdminId = ObjectId();
let UserId = ObjectId();
let TestId = ObjectId();

# 管理员所有权限
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "GET",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "POST",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "PUT",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "DELETE",
}]);


# 普通用户部分权限
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "user",
    "v1": "/user/list",
    "v2": "POST",
}]);

# 所有人均可访问（白名单）
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/posts/list",
    "v2": "GET",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/posts",
    "v2": "GET",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/user/login",
    "v2": "POST",
}]);

# 为用户授权
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "g",
    "v0": AdminId,
    "v1": "admin",
},{
    "_id": ObjectId(),
    "ptype": "g",
    "v0": UserId,
    "v1": "user",
},{
    "_id": ObjectId(),
    "ptype": "g",
    "v0": TestId,
    "v1": "test",
}]);

# 初始化用户
db.user.insertMany([{
    "_id": AdminId,
    "username": "admin",
    "password": "123456",
    "role_id": 0,
    "created_at": new Date(),
    "updated_at": new Date()
},{
    "_id": UserId,
    "username": "alice",
    "password": "123456",
    "role_id": 1,
    "created_at": new Date(),
    "updated_at": new Date()
},{
    "_id": TestId,
    "username": "test",
    "password": "123456",
    "role_id": 2,
    "created_at": new Date(),
    "updated_at": new Date()
}])

# 插入文章
db.posts.insertMany([{
    "_id": ObjectId("67b876b0e2f2b89ab22d3ca0"),
    "created_at": new Date(),
    "updated_at": new Date(),
    "title": "Stellux博客",
    "content": "## Stellux博客 & 知识库\r\n\r\n> 当你看到这篇文章的时候，代表安装成功了\r\n\r\n😀😀😀",
    "author": "codepzj",
    "like": 0
}]);

EOF
