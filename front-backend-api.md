# 备忘录后端API接口文档

## 基础信息

- **服务地址**: http://localhost:8080
- **API前缀**: /api
- **数据格式**: JSON
- **字符编码**: UTF-8

## 通用响应格式

### 成功响应
```json
{
  "code": 200,
  "message": "操作成功",
  "data": {}
}
```

### 失败响应
```json
{
  "code": 400,
  "message": "错误信息",
  "data": null
}
```

## 状态码说明

| 状态码 | 说明 |
|--------|------|
| 200 | 请求成功 |
| 400 | 请求参数错误 |
| 401 | 未授权，需要登录 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

---

## 1. 健康检查接口

### 1.1 服务状态检查

**接口地址**: `GET /health`

**请求参数**: 无

**成功响应**:
```json
{
  "status": "ok",
  "message": "备忘录后端服务运行正常",
  "version": "1.0.0"
}
```

**失败响应**:
```json
{
  "status": "error",
  "message": "服务异常"
}
```

---

## 2. 用户认证接口

### 2.1 用户登录

**接口地址**: `POST /api/auth/login`

**请求头**:
```
Content-Type: application/json
```

**请求参数**:
```json
{
  "username": "用户名",
  "password": "密码"
}
```

**参数说明**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| username | string | 是 | 用户名 |
| password | string | 是 | 用户密码 |

**成功响应**:
```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "demo_user"
    }
  }
}
```

**失败响应**:
```json
{
  "code": 400,
  "message": "用户名或密码错误",
  "data": null
}
```

### 2.2 用户注册

**接口地址**: `POST /api/auth/register`

**请求头**:
```
Content-Type: application/json
```

**请求参数**:
```json
{
  "username": "用户名",
  "password": "密码",
  "confirmPassword": "确认密码"
}
```

**参数说明**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| username | string | 是 | 用户名，3-20个字符 |
| password | string | 是 | 密码，6-20个字符 |
| confirmPassword | string | 是 | 确认密码，需与password一致 |

**成功响应**:
```json
{
  "code": 200,
  "message": "注册成功",
  "data": {
    "user": {
      "id": 2,
      "username": "new_user"
    }
  }
}
```

**失败响应**:
```json
{
  "code": 400,
  "message": "用户名已存在",
  "data": null
}
```



---

## 3. 备忘录管理接口

> **注意**: 以下接口需要在请求头中携带Authorization token
> ```
> Authorization: Bearer {token}
> ```

### 3.1 获取备忘录列表

**接口地址**: `GET /api/memos`

**请求头**:
```
Authorization: Bearer {token}
```

**请求参数** (Query参数):
| 参数名 | 类型 | 必填 | 默认值 | 说明 |
|--------|------|------|--------|------|
| page | int | 否 | 1 | 页码 |
| limit | int | 否 | 10 | 每页数量 |
| keyword | string | 否 | - | 搜索关键词 |

**成功响应**:
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "list": [
      {
        "id": 1,
        "title": "示例备忘录",
        "content": "这是一个示例备忘录内容",
        "createTime": "2024-01-01T10:00:00Z",
        "updateTime": "2024-01-01T10:00:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "limit": 10
  }
}
```

**失败响应**:
```json
{
  "code": 401,
  "message": "未授权，请先登录",
  "data": null
}
```

### 3.2 创建备忘录

**接口地址**: `POST /api/memos`

**请求头**:
```
Content-Type: application/json
Authorization: Bearer {token}
```

**请求参数**:
```json
{
  "title": "备忘录标题",
  "content": "备忘录内容"
}
```

**参数说明**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| title | string | 是 | 备忘录标题，1-100个字符 |
| content | string | 是 | 备忘录内容，1-5000个字符 |

**成功响应**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "id": 2,
    "title": "新建备忘录",
    "content": "新建的备忘录内容",
    "createTime": "2024-01-01T11:00:00Z",
    "updateTime": "2024-01-01T11:00:00Z"
  }
}
```

**失败响应**:
```json
{
  "code": 400,
  "message": "标题不能为空",
  "data": null
}
```

### 3.3 获取单个备忘录

**接口地址**: `GET /api/memos/{id}`

**请求头**:
```
Authorization: Bearer {token}
```

**路径参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 备忘录ID |

**成功响应**:
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "id": 1,
    "title": "示例备忘录",
    "content": "这是一个示例备忘录内容",
    "createTime": "2024-01-01T10:00:00Z",
    "updateTime": "2024-01-01T10:00:00Z"
  }
}
```

**失败响应**:
```json
{
  "code": 404,
  "message": "备忘录不存在",
  "data": null
}
```

### 3.4 更新备忘录

**接口地址**: `PUT /api/memos/{id}`

**请求头**:
```
Content-Type: application/json
Authorization: Bearer {token}
```

**路径参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 备忘录ID |

**请求参数**:
```json
{
  "title": "更新后的标题",
  "content": "更新后的内容"
}
```

**参数说明**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| title | string | 是 | 备忘录标题，1-100个字符 |
| content | string | 是 | 备忘录内容，1-5000个字符 |

**成功响应**:
```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "id": 1,
    "title": "更新后的备忘录",
    "content": "更新后的备忘录内容",
    "createTime": "2024-01-01T10:00:00Z",
    "updateTime": "2024-01-01T12:00:00Z"
  }
}
```

**失败响应**:
```json
{
  "code": 404,
  "message": "备忘录不存在",
  "data": null
}
```

### 3.5 删除备忘录

**接口地址**: `DELETE /api/memos/{id}`

**请求头**:
```
Authorization: Bearer {token}
```

**路径参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int | 是 | 备忘录ID |

**成功响应**:
```json
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

**失败响应**:
```json
{
  "code": 404,
  "message": "备忘录不存在",
  "data": null
}
```

---

## 4. 错误码说明

| 错误码 | 说明 | 解决方案 |
|--------|------|----------|
| 400 | 请求参数错误 | 检查请求参数格式和必填项 |
| 401 | 未授权 | 检查token是否有效，重新登录 |
| 403 | 禁止访问 | 检查用户权限 |
| 404 | 资源不存在 | 检查请求的资源ID是否正确 |
| 500 | 服务器内部错误 | 联系后端开发人员 |

---

## 5. 前端对接注意事项

1. **请求头设置**:
   - 所有POST/PUT请求需要设置 `Content-Type: application/json`
   - 需要认证的接口必须携带 `Authorization: Bearer {token}`

2. **Token管理**:
   - 登录成功后保存token到本地存储
   - 每次请求前检查token是否存在
   - 收到401错误时清除token并跳转到登录页

3. **错误处理**:
   - 统一处理HTTP状态码和业务错误码
   - 根据错误信息给用户友好的提示

4. **数据格式**:
   - 时间格式统一使用ISO 8601格式 (YYYY-MM-DDTHH:mm:ssZ)
   - 所有接口返回的数据都在data字段中

5. **分页处理**:
   - 列表接口支持分页，注意处理total、page、limit字段
   - 前端需要实现分页组件

---

## 6. 测试示例

### 使用curl测试登录接口
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "demo_user",
    "password": "123456"
  }'
```

### 使用curl测试获取备忘录列表
```bash
curl -X GET "http://localhost:8080/api/memos?page=1&limit=10" \
  -H "Authorization: Bearer your_token_here"
```

### 使用curl测试创建备忘录
```bash
curl -X POST http://localhost:8080/api/memos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_token_here" \
  -d '{
    "title": "测试备忘录",
    "content": "这是一个测试备忘录的内容"
  }'
```

---

**文档版本**: v1.0.0  
**最后更新**: 2024-01-01  
**维护人员**: 后端开发团队