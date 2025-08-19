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

// 创建算力余额集合
db.createCollection('currency_balances');

// 创建算力交易记录集合
db.createCollection('currency_transactions');

// 为用户名创建唯一索引
db.users.createIndex({ "username": 1 }, { unique: true });

// 为备忘录创建索引
db.memos.createIndex({ "user_id": 1 });
db.memos.createIndex({ "created_at": -1 });

// 为算力余额创建索引
db.currency_balances.createIndex({ "user_id": 1 }, { unique: true });
db.currency_balances.createIndex({ "last_update_time": -1 });

// 为算力交易记录创建索引
db.currency_transactions.createIndex({ "user_id": 1 });
db.currency_transactions.createIndex({ "created_at": -1 });
db.currency_transactions.createIndex({ "transaction_id": 1 }, { unique: true });
db.currency_transactions.createIndex({ "type": 1 });

// 为备忘录标题和内容创建文本搜索索引（可选，用于优化搜索性能）
db.memos.createIndex({ "title": "text", "content": "text" });

print('Database initialized successfully!');