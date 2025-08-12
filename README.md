# 备忘录应用后端

这是一个基于 Golang + Gin + MongoDB 的备忘录应用后端项目。

## 功能特性

- 用户注册、登录、忘记密码
- JWT 身份认证
- 备忘录的增删改查
- 备忘录列表分页和搜索
- CORS 跨域支持
- 密码加密存储

## 技术栈

- **Web框架**: Gin
- **数据库**: MongoDB
- **身份认证**: JWT
- **密码加密**: bcrypt
- **配置管理**: godotenv

## 项目结构

```
mjbackend/
├── config/          # 配置管理
├── controllers/     # 控制器层
├── database/        # 数据库连接
├── middleware/      # 中间件
├── models/          # 数据模型
├── routes/          # 路由配置
├── services/        # 业务逻辑层
├── utils/           # 工具函数
├── .env            # 环境配置文件
├── go.mod          # Go模块文件
├── main.go         # 程序入口
└── README.md       # 项目说明
```

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 配置环境变量

复制 `.env` 文件并根据实际情况修改配置：

```bash
# 服务器配置
PORT=8080
GIN_MODE=debug

# MongoDB配置
MONGO_URI=mongodb://admin:Password@1@192.168.22.113:30017
MONGO_DATABASE=memo_app

# JWT配置
JWT_SECRET=your-secret-key-here-change-in-production
JWT_EXPIRES_HOURS=24

# 其他配置
BCRYPT_COST=12
```

### 3. 启动服务

```bash
go run main.go
```

服务将在 `http://localhost:8080` 启动。

## API 接口

### 认证接口

- `POST /api/auth/register` - 用户注册
- `POST /api/auth/login` - 用户登录
- `POST /api/auth/forgot-password` - 忘记密码

### 备忘录接口（需要认证）

- `GET /api/memos` - 获取备忘录列表
- `POST /api/memos` - 创建备忘录
- `GET /api/memos/:id` - 获取备忘录详情
- `PUT /api/memos/:id` - 更新备忘录
- `DELETE /api/memos/:id` - 删除备忘录

### 其他接口

- `GET /health` - 健康检查

## 数据库配置

### MongoDB 连接信息

- **主机**: 192.168.22.113
- **端口**: 30017
- **用户名**: admin
- **密码**: Password@1
- **数据库**: memo_app

### 集合结构

#### users 集合
```json
{
  "_id": "ObjectId",
  "username": "string",
  "phone": "string",
  "password": "string (encrypted)",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

#### memos 集合
```json
{
  "_id": "ObjectId",
  "user_id": "ObjectId",
  "title": "string",
  "content": "string",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

## 开发说明

### 身份认证

所有备忘录相关的接口都需要在请求头中携带 JWT token：

```
Authorization: Bearer <your-jwt-token>
```

### 错误处理

API 返回统一的错误格式：

```json
{
  "code": 400,
  "message": "错误信息"
}
```

### 成功响应

API 返回统一的成功格式：

```json
{
  "code": 200,
  "message": "操作成功",
  "data": {}
}
```

## 部署说明

1. 修改 `.env` 文件中的配置
2. 设置 `GIN_MODE=release`
3. 更改 `JWT_SECRET` 为安全的密钥
4. 构建并运行：

```bash
go build -o mjbackend
./mjbackend
```

## 注意事项

1. 生产环境请务必更改 JWT 密钥
2. 建议使用 HTTPS
3. 定期备份数据库
4. 监控服务运行状态# mjapp
