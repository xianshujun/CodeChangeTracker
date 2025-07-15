# API 设计文档

## 概述

CodeChangeTracker API 提供了一套完整的RESTful接口，用于管理代码变更风险分析系统。

## 基础信息

- **API版本**: v1
- **基础URL**: `http://localhost:8080/api/v1`
- **认证方式**: JWT Token
- **数据格式**: JSON

## 认证

所有API请求都需要在请求头中包含JWT令牌：

```
Authorization: Bearer <your-jwt-token>
```

## 错误处理

API使用标准HTTP状态码来表示请求结果：

- `200 OK` - 成功
- `201 Created` - 创建成功
- `400 Bad Request` - 请求参数错误
- `401 Unauthorized` - 未授权
- `403 Forbidden` - 禁止访问
- `404 Not Found` - 资源不存在
- `500 Internal Server Error` - 服务器内部错误

错误响应格式：
```json
{
  "error": "错误描述",
  "code": "ERROR_CODE",
  "message": "详细错误信息"
}
```

## 项目管理 API

### 创建项目

```http
POST /api/v1/projects
Content-Type: application/json

{
  "name": "项目名称",
  "description": "项目描述",
  "repo_url": "https://github.com/user/repo.git",
  "language": "go",
  "branch": "main"
}
```

响应：
```json
{
  "id": 1,
  "name": "项目名称",
  "description": "项目描述",
  "repo_url": "https://github.com/user/repo.git",
  "language": "go",
  "branch": "main",
  "is_active": true,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### 获取项目列表

```http
GET /api/v1/projects?page=1&limit=10&language=go
```

查询参数：
- `page` (可选): 页码，默认为1
- `limit` (可选): 每页数量，默认为10，最大100
- `language` (可选): 编程语言过滤
- `is_active` (可选): 项目状态过滤

响应：
```json
{
  "data": [
    {
      "id": 1,
      "name": "项目名称",
      "description": "项目描述",
      "repo_url": "https://github.com/user/repo.git",
      "language": "go",
      "created_at": "2023-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 50,
    "total_pages": 5
  }
}
```

### 获取项目详情

```http
GET /api/v1/projects/{id}
```

### 更新项目

```http
PUT /api/v1/projects/{id}
Content-Type: application/json

{
  "name": "更新的项目名称",
  "description": "更新的项目描述"
}
```

### 删除项目

```http
DELETE /api/v1/projects/{id}
```

### 触发项目分析

```http
POST /api/v1/projects/{id}/analyze
```

响应：
```json
{
  "message": "分析任务已启动",
  "task_id": "uuid-task-id"
}
```

## 代码变更 API

### 分析代码变更

```http
POST /api/v1/changes/analyze
Content-Type: application/json

{
  "project_id": 1,
  "commit_hash": "abc123",
  "author": "developer@example.com",
  "message": "fix: 修复bug",
  "files_changed": ["src/main.go", "src/utils.go"],
  "lines_added": 10,
  "lines_deleted": 5
}
```

响应：
```json
{
  "id": 1,
  "project_id": 1,
  "commit_hash": "abc123",
  "risk_score": 0.75,
  "status": "analyzed",
  "created_at": "2023-01-01T00:00:00Z"
}
```

### 获取变更详情

```http
GET /api/v1/changes/{id}
```

### 获取风险评估结果

```http
GET /api/v1/changes/{id}/assessment
```

响应：
```json
{
  "id": 1,
  "change_id": 1,
  "risk_level": "MEDIUM",
  "score": 0.75,
  "factors": [
    "修改了核心业务逻辑",
    "涉及数据库操作",
    "缺少单元测试"
  ],
  "suggestions": [
    "建议增加单元测试",
    "建议进行代码审查",
    "建议在测试环境充分验证"
  ],
  "confidence": 0.85,
  "model_version": "v1.0.0",
  "created_at": "2023-01-01T00:00:00Z"
}
```

## 分析报告 API

### 生成项目报告

```http
GET /api/v1/analysis/reports/{project_id}?from=2023-01-01&to=2023-12-31&format=json
```

查询参数：
- `from` (可选): 开始日期，格式：YYYY-MM-DD
- `to` (可选): 结束日期，格式：YYYY-MM-DD
- `format` (可选): 报告格式，支持json、pdf、excel

### 获取项目度量指标

```http
GET /api/v1/analysis/metrics/{project_id}?period=30d
```

查询参数：
- `period` (可选): 时间周期，支持7d、30d、90d、1y

响应：
```json
{
  "project_id": 1,
  "period": "30d",
  "metrics": {
    "total_changes": 100,
    "high_risk_changes": 5,
    "medium_risk_changes": 20,
    "low_risk_changes": 75,
    "average_risk_score": 0.35,
    "trend": {
      "risk_score_trend": "decreasing",
      "change_frequency": "stable"
    }
  },
  "generated_at": "2023-01-01T00:00:00Z"
}
```

## Webhook

### 配置Webhook

系统支持通过Webhook接收Git仓库的推送事件：

```http
POST /api/v1/webhooks/git
Content-Type: application/json
X-Hub-Signature-256: sha256=...

{
  "repository": {
    "full_name": "user/repo"
  },
  "commits": [
    {
      "id": "abc123",
      "author": {
        "email": "developer@example.com"
      },
      "message": "fix: 修复bug",
      "added": ["src/new.go"],
      "modified": ["src/main.go"],
      "removed": []
    }
  ]
}
```

## 限流

API实施了限流机制：
- 每个IP每分钟最多100个请求
- 认证用户每分钟最多1000个请求
- 分析接口每分钟最多10个请求

当达到限流阈值时，会返回429状态码：

```json
{
  "error": "Rate limit exceeded",
  "retry_after": 60
}
```