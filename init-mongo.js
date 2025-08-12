// MongoDB 初始化脚本
// 创建应用数据库和用户

db = db.getSiblingDB('mjbackend');

// 创建应用用户
db.createUser({
  user: 'mjuser',
  pwd: 'mjpass@123',
  roles: [
    {
      role: 'readWrite',
      db: 'mjbackend'
    }
  ]
});

// 创建用户集合
db.createCollection('users');

// 创建备忘录集合
db.createCollection('memos');

// 为用户名和手机号创建唯一索引
db.users.createIndex({ "username": 1 }, { unique: true });
db.users.createIndex({ "phone": 1 }, { unique: true });

// 为备忘录创建索引
db.memos.createIndex({ "user_id": 1 });
db.memos.createIndex({ "created_at": -1 });

print('Database initialized successfully!');