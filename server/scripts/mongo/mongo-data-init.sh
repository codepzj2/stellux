mongosh -- "$MONGO_INITDB_DATABASE" <<EOF
db = db.getSiblingDB('$MONGO_INITDB_DATABASE')
db.auth('$MONGO_USERNAME', '$MONGO_PASSWORD');


# 用户表为用户名创建唯一升序索引
db.user.createIndex({"username":1},{"unique":true});

let AdminId = ObjectId();
let StaffId = ObjectId();

# 初始化权限表
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "*",
    "v2": "GET",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/user/login",
    "v2": "POST",
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "/posts/create",
    "v2": "POST",
},{
    "_id": ObjectId(),
    "ptype": "g",
    "v0": AdminId,
    "v1": "admin",
},{
    "_id": ObjectId(),
    "ptype": "g",
    "v0": StaffId,
    "v1": "staff",
},
])

# 初始化用户
db.user.insertMany([{
    "_id": AdminId,
    "username": "admin",
    "password": "123456",
    "role_id": 0,
    "created_at": new Date(),
    "updated_at": new Date()
},{
    "_id": StaffId,
    "username": "staff",
    "password": "123456",
    "role_id": 1,
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
