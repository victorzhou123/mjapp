# 备忘录应用 API 接口文档

这是一个基于golang + gin + gorm + mongodb的备忘录的后端项目，mongodb的地址为`mongodb://192.168.22.113:30017`，用户名为admin，密码为Password@1

## 概述
本文档定义了备忘录应用前端所需的后端API接口，包括用户认证和备忘录管理功能。

## 基础信息
- **Base URL**: `https://api.yourdomain.com`
- **Content-Type**: `application/json`
- **认证方式**: Bearer Token

## 1. 用户认证模块

### 1.1 用户登录
**接口地址**: `POST /api/auth/login`

**请求参数**:
```json
{
  "username": "string",  // 用户名或手机号
  "password": "string"   // 密码
}
```

**响应数据**:
```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "token": "string",     // JWT token
    "user": {
      "id": "number",
      "username": "string",
      "phone": "string"
    }
  }
}
```

**错误响应**:
```json
{
  "code": 400,
  "message": "用户名或密码错误"
}
```

### 1.2 用户注册
**接口地址**: `POST /api/auth/register`

**请求参数**:
```json
{
  "username": "string",     // 用户名
  "phone": "string",        // 手机号
  "password": "string",     // 密码
  "confirmPassword": "string" // 确认密码
}
```

**响应数据**:
```json
{
  "code": 200,
  "message": "注册成功",
  "data": {
    "user": {
      "id": "number",
      "username": "string",
      "phone": "string"
    }
  }
}
```

**错误响应**:
```json
{
  "code": 400,
  "message": "用户名已存在" // 或其他错误信息
}
```

### 1.3 忘记密码
**接口地址**: `POST /api/auth/forgot-password`

**请求参数**:
```json
{
  "phone": "string"  // 手机号
}
```

**响应数据**:
```json
{
  "code": 200,
  "message": "验证码已发送"
}
```

## 2. 备忘录管理模块

### 2.1 获取备忘录列表
**接口地址**: `GET /api/memos`

**请求头**:
```
Authorization: Bearer {token}
```

**查询参数**:
- `page`: 页码 (可选，默认1)
- `limit`: 每页数量 (可选，默认10)
- `keyword`: 搜索关键词 (可选)

**响应数据**:
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "list": [
      {
        "id": "number",
        "title": "string",
        "content": "string",
        "createTime": "string",  // ISO 8601 格式
        "updateTime": "string"
      }
    ],
    "total": "number",
    "page": "number",
    "limit": "number"
  }
}
```

### 2.2 获取备忘录详情
**接口地址**: `GET /api/memos/{id}`

**请求头**:
```
Authorization: Bearer {token}
```

**路径参数**:
- `id`: 备忘录ID

**响应数据**:
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "id": "number",
    "title": "string",
    "content": "string",
    "createTime": "string",
    "updateTime": "string"
  }
}
```

### 2.3 创建备忘录
**接口地址**: `POST /api/memos`

**请求头**:
```
Authorization: Bearer {token}
```

**请求参数**:
```json
{
  "title": "string",    // 标题，必填
  "content": "string"  // 内容，可选
}
```

**响应数据**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "id": "number",
    "title": "string",
    "content": "string",
    "createTime": "string",
    "updateTime": "string"
  }
}
```

### 2.4 更新备忘录
**接口地址**: `PUT /api/memos/{id}`

**请求头**:
```
Authorization: Bearer {token}
```

**路径参数**:
- `id`: 备忘录ID

**请求参数**:
```json
{
  "title": "string",    // 标题，必填
  "content": "string"  // 内容，可选
}
```

**响应数据**:
```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "id": "number",
    "title": "string",
    "content": "string",
    "createTime": "string",
    "updateTime": "string"
  }
}
```

### 2.5 删除备忘录
**接口地址**: `DELETE /api/memos/{id}`

**请求头**:
```
Authorization: Bearer {token}
```

**路径参数**:
- `id`: 备忘录ID

**响应数据**:
```json
{
  "code": 200,
  "message": "删除成功"
}
```

## 3. 通用错误码

| 错误码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权，token无效或过期 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

## 4. 数据库设计建议

### 用户表 (users)
```sql
CREATE TABLE users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(50) UNIQUE NOT NULL,
  phone VARCHAR(20) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### 备忘录表 (memos)
```sql
CREATE TABLE memos (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  title VARCHAR(255) NOT NULL,
  content TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
```

## 5. 开发注意事项

1. **安全性**:
   - 密码需要加密存储（建议使用bcrypt）
   - JWT token设置合理的过期时间
   - 对用户输入进行验证和过滤

2. **性能优化**:
   - 备忘录列表支持分页
   - 添加必要的数据库索引
   - 考虑缓存热点数据

3. **数据验证**:
   - 用户名长度限制：3-20字符
   - 手机号格式验证
   - 密码强度要求：至少6位
   - 备忘录标题不能为空

4. **跨域处理**:
   - 配置CORS允许前端域名访问

5. **日志记录**:
   - 记录关键操作日志
   - 记录错误信息便于调试