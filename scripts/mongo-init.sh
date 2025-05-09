#!/bin/bash

mongosh -u "$MONGO_INITDB_ROOT_USERNAME" -p "$MONGO_INITDB_ROOT_PASSWORD" --authenticationDatabase admin <<EOF

// 创建目标数据库用户
db = db.getSiblingDB('$MONGO_INITDB_DATABASE');
db.createUser({
    user: '$MONGO_USERNAME',
    pwd: '$MONGO_PASSWORD',
    roles: [{ role: 'readWrite', db: '$MONGO_INITDB_DATABASE' }]
});

db.auth('$MONGO_USERNAME', '$MONGO_PASSWORD');

// 用户表为用户名创建唯一升序索引
db.user.createIndex({"username": 1}, {"unique": true});

// 文章表创建文本索引
db.post.createIndex({
    "title": "text",
    "content": "text",
    "description": "text"
});

let AdminId = ObjectId("67c453eda04b00c407b431fd");
let UserId = ObjectId("67c453eda04b00c407b431fe");
let TestId = ObjectId("67c453eda04b00c407b431ff");

// 管理员所有权限
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "GET"
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "POST"
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "PATCH"
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "PUT"
},{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "admin",
    "v1": "*",
    "v2": "DELETE"
}]);

// 普通用户权限
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "user",
    "v1": "/admin-api/user/list*",
    "v2": "GET"
}]);

// 测试用户权限
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/admin-api/user/info",
    "v2": "GET"
}, {
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/posts/*",
    "v2": "GET"
}, {
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/user/login",
    "v2": "POST"
}, {
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/images/*",
    "v2": "GET"
}, {
    "_id": ObjectId(),
    "ptype": "p",
    "v0": "*",
    "v1": "/picture/list*",
    "v2": "GET"
}]);

// 为用户授权
db.casbin_rule.insertMany([{
    "_id": ObjectId(),
    "ptype": "g",
    "v0": "67c453eda04b00c407b431fd",
    "v1": "admin"
}, {
    "_id": ObjectId(),
    "ptype": "g",
    "v0": "67c453eda04b00c407b431fe",
    "v1": "user"
}, {
    "_id": ObjectId(),
    "ptype": "g",
    "v0": "67c453eda04b00c407b431ff",
    "v1": "test"
}]);

// 初始化用户
db.user.insertMany([{
    "_id": AdminId,
    "username": "admin",
    "password": "\$2a\$10\$SLcnDmaJc1nLtUOsZS4yquXyVeu5E6qJHNTVeKSzTk4JO4Xq/FPSy",
    "nickname": "芒果",
    "role_id": 0,
    "created_at": new Date(),
    "updated_at": new Date(),
    "avatar": "https://github.githubassets.com/assets/pull-shark-default-498c279a747d.png",
    "email": "admin@example.com",
}, {
    "_id": UserId,
    "username": "alice",
    "password": "\$2a\$10\$SLcnDmaJc1nLtUOsZS4yquXyVeu5E6qJHNTVeKSzTk4JO4Xq/FPSy",
    "nickname": "小李",
    "role_id": 1,
    "created_at": new Date(),
    "updated_at": new Date(),
    "avatar": "https://github.githubassets.com/assets/quickdraw-default-39c6aec8ff89.png",
    "email": "alice@example.com",
}, {
    "_id": TestId,
    "username": "test",
    "password": "\$2a\$10\$SLcnDmaJc1nLtUOsZS4yquXyVeu5E6qJHNTVeKSzTk4JO4Xq/FPSy",
    "nickname": "小张",
    "role_id": 2,
    "created_at": new Date(),
    "updated_at": new Date(),
    "avatar": "https://github.githubassets.com/assets/yolo-default-be0bbff04951.png",
    "email": "test@example.com",
}]);

// 初始化分类
db.label.insertMany([{
    "_id": ObjectId("67c453eda04b00c407b43197"),
    "label_type": "category",
    "name": "技术",
}, {
    "_id": ObjectId("67c453eda04b00c407b43198"),
    "label_type": "category",
    "name": "生活",
}]);

// 初始化标签
db.label.insertMany([{
    "_id": ObjectId("67c453eda04b00c407b43199"),
    "label_type": "tag",
    "name": "golang",
}, {
    "_id": ObjectId("67c453eda04b00c407b43200"),
    "label_type": "tag",
    "name": "python",
}, {
    "_id": ObjectId("67c453eda04b00c407b43201"),
    "label_type": "tag",
    "name": "java",
}]);

// 初始化文章
db.post.insertMany([{
    "_id": ObjectId("67c453eda04b00c407b43202"),
    "created_at": new Date(),
    "updated_at": new Date(),
    "title": "stellux博客?",
    "content": "STELLUX 是一个基于 Go 语言的博客系统，它可以帮助你快速搭建一个博客，并且支持 Markdown 格式\n\n## 什么是 STELLUX?\n\nSTELLUX 是一个基于 Go 语言的博客系统，它可以帮助你快速搭建一个博客，并且支持 Markdown 格式。\n\n## 为什么选择 STELLUX?\n\nSTELLUX 是一个基于 Go 语言的博客系统，它可以帮助你快速搭建一个博客，并且支持 Markdown 格式。\n\n## 如何使用 STELLUX?\n\nSTELLUX 是一个基于 Go 语言的博客系统，它可以帮助你快速搭建一个博客，并且支持 Markdown 格式。",
    "description": "STELLUX 是一个基于 Go 语言的博客系统，它可以帮助你快速搭建一个博客，并且支持 Markdown 格式。",
    "category_id": ObjectId("67c453eda04b00c407b43197"),
    "tags_id": [ObjectId("67c453eda04b00c407b43199"), ObjectId("67c453eda04b00c407b43201")],
}]);

EOF