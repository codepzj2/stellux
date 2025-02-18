mongosh -- "$MONGO_INITDB_DATABASE" <<EOF
db = db.getSiblingDB('$MONGO_INITDB_DATABASE')
db.auth('$MONGO_USERNAME', '$MONGO_PASSWORD');

db.user.insertOne({
    "username": "admin",
    "password": "123456",
    "role_id": 0,
    "created_at": new Date(),
    "updated_at": new Date()
});

EOF