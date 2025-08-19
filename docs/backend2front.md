# 备忘录系统前端接口文档

## 基础信息

- **服务地址**: http://localhost:8080
- **API前缀**: /api
- **数据格式**: JSON
- **字符编码**: UTF-8
- **文档版本**: v1.0.0
- **最后更新**: 2024-01-01

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

### 带数据的错误响应
```json
{
  "code": 400,
  "message": "错误信息",
  "data": {
    "currentBalance": 50,
    "requiredAmount": 100
  }
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
      "id": "507f1f77bcf86cd799439011",
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
      "id": "507f1f77bcf86cd799439012",
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
| page | int | 否 | 1 | 页码，从1开始 |
| limit | int | 否 | 10 | 每页数量，最大100 |
| keyword | string | 否 | - | 搜索关键词，支持标题和内容搜索 |

**成功响应**:
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "list": [
      {
        "id": "507f1f77bcf86cd799439011",
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
  "message": "用户未认证",
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
| title | string | 是 | 备忘录标题 |
| content | string | 否 | 备忘录内容 |

**成功响应**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "title": "备忘录标题",
    "content": "备忘录内容",
    "createTime": "2024-01-01T10:00:00Z",
    "updateTime": "2024-01-01T10:00:00Z"
  }
}
```

**失败响应**:
```json
{
  "code": 400,
  "message": "请求参数错误",
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
| id | string | 是 | 备忘录ID (ObjectID格式) |

**成功响应**:
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "id": "507f1f77bcf86cd799439011",
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
| id | string | 是 | 备忘录ID (ObjectID格式) |

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
| title | string | 是 | 备忘录标题 |
| content | string | 否 | 备忘录内容 |

**成功响应**:
```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "id": "507f1f77bcf86cd799439011",
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
| id | string | 是 | 备忘录ID (ObjectID格式) |

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

## 4. 算力管理接口

> **注意**: 以下接口需要在请求头中携带Authorization token
> ```
> Authorization: Bearer {token}
> ```

### 4.1 查询算力余额

**接口地址**: `GET /api/currency/balance`

**请求头**:
```
Authorization: Bearer {token}
```

**请求参数**: 无

**成功响应**:
```json
{
  "code": 200,
  "message": "查询成功",
  "data": {
    "balance": 100,
    "lastUpdateTime": "2024-01-01T12:00:00Z"
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

### 4.2 扣减算力

**接口地址**: `POST /api/currency/deduct`

**请求头**:
```
Content-Type: application/json
Authorization: Bearer {token}
```

**请求参数**:
```json
{
  "amount": 10,
  "reason": "创建备忘录",
  "memoId": "507f1f77bcf86cd799439011"
}
```

**参数说明**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| amount | int | 是 | 扣减数量，必须大于0 |
| reason | string | 是 | 扣减原因 |
| memoId | string | 否 | 关联的备忘录ID (ObjectID格式) |

**成功响应**:
```json
{
  "code": 200,
  "message": "扣减成功",
  "data": {
    "remainingBalance": 90,
    "deductedAmount": 10,
    "transactionId": "txn_1234567890"
  }
}
```

**余额不足响应**:
```json
{
  "code": 400,
  "message": "算力余额不足",
  "data": {
    "currentBalance": 5,
    "requiredAmount": 10
  }
}
```

**失败响应**:
```json
{
  "code": 500,
  "message": "扣减算力失败: 系统错误",
  "data": null
}
```

### 4.3 充值算力

**接口地址**: `POST /api/currency/recharge`

**请求头**:
```
Content-Type: application/json
Authorization: Bearer {token}
```

**请求参数**:
```json
{
  "amount": 100,
  "transactionId": "pay_1234567890",
  "source": "支付宝"
}
```

**参数说明**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| amount | int | 是 | 充值数量，必须大于0 |
| transactionId | string | 是 | 支付交易ID，用于防重复充值 |
| source | string | 否 | 充值来源（如：支付宝、微信等） |

**成功响应**:
```json
{
  "code": 200,
  "message": "充值成功",
  "data": {
    "newBalance": 200,
    "rechargedAmount": 100,
    "transactionId": "pay_1234567890"
  }
}
```

**重复充值响应**:
```json
{
  "code": 400,
  "message": "交易ID已存在，请勿重复充值",
  "data": null
}
```

**失败响应**:
```json
{
  "code": 400,
  "message": "支付凭证验证失败",
  "data": null
}
```

---

## 5. 数据类型说明

### 5.1 ObjectID格式
- MongoDB的ObjectID是24位十六进制字符串
- 示例: `507f1f77bcf86cd799439011`
- 前端在处理ID时应当作字符串处理

### 5.2 时间格式
- 统一使用ISO 8601格式: `YYYY-MM-DDTHH:mm:ssZ`
- 示例: `2024-01-01T10:00:00Z`
- 时区为UTC

### 5.3 分页数据结构
```json
{
  "list": [],      // 数据列表
  "total": 100,    // 总记录数
  "page": 1,       // 当前页码
  "limit": 10      // 每页数量
}
```

---

## 6. 错误处理指南

### 6.1 通用错误处理

| 错误码 | 说明 | 前端处理建议 |
|--------|------|-------------|
| 400 | 请求参数错误 | 检查请求参数格式和必填项，显示具体错误信息 |
| 401 | 未授权 | 清除本地token，跳转到登录页 |
| 403 | 禁止访问 | 显示权限不足提示 |
| 404 | 资源不存在 | 显示资源不存在提示，可能需要刷新列表 |
| 500 | 服务器内部错误 | 显示系统繁忙提示，建议稍后重试 |

### 6.2 特殊错误处理

#### 算力余额不足
```json
{
  "code": 400,
  "message": "算力余额不足",
  "data": {
    "currentBalance": 5,
    "requiredAmount": 10
  }
}
```
前端可以根据`data`中的信息提示用户当前余额和所需金额，引导用户充值。

#### 重复操作
```json
{
  "code": 400,
  "message": "交易ID已存在，请勿重复充值",
  "data": null
}
```
前端应该防止用户重复提交，可以在提交后禁用按钮或显示loading状态。

---

## 7. 前端对接注意事项

### 7.1 请求头设置
- 所有POST/PUT请求需要设置 `Content-Type: application/json`
- 需要认证的接口必须携带 `Authorization: Bearer {token}`
- token格式: `Bearer ` + JWT字符串（注意Bearer后有一个空格）

### 7.2 Token管理
- 登录成功后保存token到本地存储（localStorage或sessionStorage）
- 每次请求前检查token是否存在
- 收到401错误时清除token并跳转到登录页
- 建议实现token自动刷新机制

### 7.3 请求拦截器建议
```javascript
// 请求拦截器
axios.interceptors.request.use(config => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// 响应拦截器
axios.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);
```

### 7.4 分页处理
- 列表接口支持分页，注意处理total、page、limit字段
- 前端需要实现分页组件
- 建议实现无限滚动或传统分页两种模式

### 7.5 搜索功能
- 备忘录列表支持关键词搜索
- 搜索会匹配标题和内容
- 建议实现防抖功能，避免频繁请求

### 7.6 数据验证
- 前端应该实现客户端验证，提升用户体验
- 用户名：3-20个字符
- 密码：6-20个字符
- 备忘录标题：必填
- 算力数量：必须大于0的整数

---

## 8. 测试示例

### 8.1 使用curl测试登录接口
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "demo_user",
    "password": "123456"
  }'
```

### 8.2 使用curl测试获取备忘录列表
```bash
curl -X GET "http://localhost:8080/api/memos?page=1&limit=10" \
  -H "Authorization: Bearer your_token_here"
```

### 8.3 使用curl测试创建备忘录
```bash
curl -X POST http://localhost:8080/api/memos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_token_here" \
  -d '{
    "title": "测试备忘录",
    "content": "这是一个测试备忘录的内容"
  }'
```

### 8.4 使用curl测试查询算力余额
```bash
curl -X GET http://localhost:8080/api/currency/balance \
  -H "Authorization: Bearer your_token_here"
```

### 8.5 使用curl测试扣减算力
```bash
curl -X POST http://localhost:8080/api/currency/deduct \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_token_here" \
  -d '{
    "amount": 10,
    "reason": "创建备忘录",
    "memoId": "507f1f77bcf86cd799439011"
  }'
```

### 8.6 使用curl测试充值算力
```bash
curl -X POST http://localhost:8080/api/currency/recharge \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_token_here" \
  -d '{
    "amount": 100,
    "transactionId": "pay_1234567890",
    "source": "支付宝"
  }'
```

---

## 9. 接口清单

### 9.1 无需认证的接口
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /health | 健康检查 |
| POST | /api/auth/login | 用户登录 |
| POST | /api/auth/register | 用户注册 |

### 9.2 需要认证的接口
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/memos | 获取备忘录列表 |
| POST | /api/memos | 创建备忘录 |
| GET | /api/memos/{id} | 获取单个备忘录 |
| PUT | /api/memos/{id} | 更新备忘录 |
| DELETE | /api/memos/{id} | 删除备忘录 |
| GET | /api/currency/balance | 查询算力余额 |
| POST | /api/currency/deduct | 扣减算力 |
| POST | /api/currency/recharge | 充值算力 |

---

## 10. 常见问题

### Q1: 如何处理token过期？
A1: 当收到401状态码时，说明token已过期或无效，前端应该清除本地token并跳转到登录页。

### Q2: 备忘录ID的格式是什么？
A2: 备忘录ID是MongoDB的ObjectID，24位十六进制字符串，如：`507f1f77bcf86cd799439011`

### Q3: 分页参数有什么限制？
A3: page最小为1，limit最小为1最大为100，超出范围会使用默认值。

### Q4: 算力扣减失败如何处理？
A4: 如果是余额不足，会返回特殊的错误响应，包含当前余额和所需金额信息，前端可以据此引导用户充值。

### Q5: 如何防止重复充值？
A5: 每次充值都需要提供唯一的transactionId，系统会检查该ID是否已存在，如果存在则拒绝充值。

### Q6: 搜索功能支持哪些字段？
A6: 目前支持备忘录的标题和内容搜索，使用keyword参数。

---

**文档维护**: 后端开发团队  
**技术支持**: 如有问题请联系后端开发人员