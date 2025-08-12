# Docker 部署指南

本项目已配置 Docker 容器化部署，支持本地开发和生产环境部署。

## 文件说明

- `Dockerfile`: 应用容器构建文件
- `docker-compose.yml`: 容器编排配置
- `.dockerignore`: Docker 构建忽略文件
- `init-mongo.js`: MongoDB 初始化脚本

## 快速开始

### 1. 使用 Docker Compose 启动（推荐）

```bash
# 构建并启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f mjbackend
```

### 2. 单独构建应用镜像

```bash
# 构建镜像
docker build -t mjbackend:latest .

# 运行容器（需要先启动 MongoDB）
docker run -d \
  --name mjbackend-app \
  -p 8080:8080 \
  -e MONGO_URI="mongodb://mjuser:mjpass%40123@mongodb:27017/mjbackend?authSource=mjbackend" \
  -e JWT_SECRET="your-jwt-secret" \
  mjbackend:latest
```

## 环境配置

### 环境变量

应用支持以下环境变量：

- `GIN_MODE`: Gin 运行模式 (debug/release)
- `MONGO_URI`: MongoDB 连接字符串
- `JWT_SECRET`: JWT 签名密钥
- `PORT`: 应用端口 (默认 8080)

### MongoDB 配置

Docker Compose 会自动创建：
- 数据库: `mjbackend`
- 用户: `mjuser`
- 密码: `mjpass@123`

## 服务访问

启动成功后，可以通过以下地址访问：

- 应用服务: http://localhost:8080
- 健康检查: http://localhost:8080/health
- MongoDB: localhost:27017

## 常用命令

```bash
# 停止所有服务
docker-compose down

# 停止并删除数据卷
docker-compose down -v

# 重新构建并启动
docker-compose up --build -d

# 进入应用容器
docker-compose exec mjbackend sh

# 进入 MongoDB 容器
docker-compose exec mongodb mongosh

# 查看实时日志
docker-compose logs -f
```

## 生产环境部署

### 1. 修改配置

生产环境部署前，请修改以下配置：

1. 更改 `docker-compose.yml` 中的数据库密码
2. 设置强密码的 JWT_SECRET
3. 配置适当的资源限制
4. 启用 HTTPS

### 2. 安全建议

- 使用强密码
- 限制数据库访问
- 配置防火墙规则
- 定期备份数据
- 使用 HTTPS

### 3. 监控和日志

```bash
# 查看容器资源使用情况
docker stats

# 导出日志
docker-compose logs mjbackend > app.log
```

## 故障排除

### 常见问题

1. **端口冲突**
   - 修改 `docker-compose.yml` 中的端口映射

2. **MongoDB 连接失败**
   - 检查 MongoDB 容器是否正常启动
   - 验证连接字符串格式

3. **应用启动失败**
   - 查看应用日志: `docker-compose logs mjbackend`
   - 检查环境变量配置

### 数据备份

```bash
# 备份 MongoDB 数据
docker-compose exec mongodb mongodump --db mjbackend --out /data/backup

# 恢复数据
docker-compose exec mongodb mongorestore /data/backup
```

## 开发环境

开发时可以使用卷挂载实现热重载：

```yaml
volumes:
  - .:/app
  - /app/vendor
```

注意：生产环境不建议使用卷挂载。