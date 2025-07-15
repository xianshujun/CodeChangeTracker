# CodeChangeTracker

<div align="center">
  <h3>🚀 AI-powered Code Change Risk Analysis</h3>
  <p>智能代码变更风险分析系统</p>
  
  [![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
  [![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
  [![API](https://img.shields.io/badge/API-REST-orange.svg)](docs/api.md)
  [![Architecture](https://img.shields.io/badge/Architecture-Microservices-purple.svg)](docs/architecture.md)
</div>

## 📖 项目概述

CodeChangeTracker 是一个基于人工智能的代码变更风险分析系统，旨在帮助开发团队：

- 🔍 **智能分析**: 使用AI技术分析代码变更的潜在风险
- 📊 **风险评估**: 提供详细的风险评分和改进建议
- 📈 **趋势追踪**: 跟踪项目代码质量和风险趋势
- 🔄 **CI/CD集成**: 无缝集成到现有开发工作流
- 📋 **详细报告**: 生成可视化分析报告

## ✨ 核心功能

### 🔬 代码分析
- 多语言支持 (Go, Java, Python, JavaScript, 等)
- 静态代码分析
- 复杂度评估
- 代码覆盖率分析

### 🤖 AI 风险评估
- 基于机器学习的风险预测
- 历史数据学习
- 智能建议生成
- 持续模型优化

### 📊 可视化报告
- 实时风险仪表盘
- 趋势分析图表
- PDF/Excel 报告导出
- 团队协作界面

### 🔧 集成支持
- Git Webhooks
- CI/CD 管道集成
- IDE 插件支持
- API 接口开放

## 🏗️ 系统架构

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   客户端层       │    │   API网关层      │    │   应用服务层     │
│  Web UI/CLI     │ ←→ │  认证/路由/限流   │ ←→ │  业务逻辑处理     │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                                        ↓
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   AI分析引擎     │    │   数据存储层     │    │   外部集成层     │
│  风险评估/ML     │ ←→ │ PostgreSQL/Redis │ ←→ │  Git/CI/CD      │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

详细架构设计请参考：[Architecture Documentation](docs/architecture.md)

## 🚀 快速开始

### 环境要求

- Go 1.21+
- PostgreSQL 15+
- Redis 7+
- Docker & Docker Compose (可选)

### 1. 克隆仓库

```bash
git clone https://github.com/xianshujun/CodeChangeTracker.git
cd CodeChangeTracker
```

### 2. 配置环境

```bash
# 复制配置文件
cp .env.example .env

# 编辑配置文件
vim .env
```

### 3. 启动服务

#### 使用 Docker Compose (推荐)

```bash
docker-compose up -d
```

#### 手动启动

```bash
# 安装依赖
go mod download

# 启动数据库服务 (如果使用Docker)
docker-compose up -d db redis

# 运行应用
go run cmd/server/main.go
```

### 4. 验证安装

```bash
# 检查服务状态
curl http://localhost:8080/health

# 访问API文档
open http://localhost:8080/swagger/index.html
```

## 📚 文档

- [架构设计](docs/architecture.md) - 系统架构和设计理念
- [API文档](docs/api.md) - RESTful API接口说明
- [开发指南](docs/development.md) - 开发环境设置和开发流程
- [部署指南](docs/deployment.md) - 生产环境部署说明

## 🛠️ 开发指南

### 项目结构

```
CodeChangeTracker/
├── cmd/                    # 应用入口
├── internal/               # 私有代码
│   ├── api/               # API层
│   ├── config/            # 配置管理
│   ├── database/          # 数据库
│   ├── models/            # 数据模型
│   ├── services/          # 业务逻辑
│   └── analyzers/         # 代码分析器
├── pkg/                   # 可重用代码
├── docs/                  # 文档
├── test/                  # 测试
└── deployments/           # 部署配置
```

### 本地开发

```bash
# 运行测试
go test ./...

# 代码格式化
go fmt ./...

# 静态分析
golint ./...

# 构建应用
go build -o bin/server cmd/server/main.go
```

更多详细信息请参考：[Development Guide](docs/development.md)

## 🔧 配置说明

### 环境变量

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| `ENVIRONMENT` | 运行环境 | `development` |
| `PORT` | 服务端口 | `8080` |
| `DATABASE_URL` | 数据库连接字符串 | - |
| `REDIS_URL` | Redis连接字符串 | - |
| `JWT_SECRET` | JWT密钥 | - |
| `OPENAI_API_KEY` | OpenAI API密钥 | - |

### 数据库配置

```sql
-- PostgreSQL 数据库设置
CREATE DATABASE codechange_tracker;
CREATE USER cct_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE codechange_tracker TO cct_user;
```

## 🌐 API 使用示例

### 创建项目

```bash
curl -X POST http://localhost:8080/api/v1/projects \
  -H "Content-Type: application/json" \
  -d '{
    "name": "示例项目",
    "repo_url": "https://github.com/user/repo.git",
    "language": "go"
  }'
```

### 分析代码变更

```bash
curl -X POST http://localhost:8080/api/v1/changes/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
    "commit_hash": "abc123",
    "author": "developer@example.com",
    "files_changed": ["main.go", "utils.go"]
  }'
```

更多API示例请参考：[API Documentation](docs/api.md)

## 🧪 测试

```bash
# 运行单元测试
go test ./...

# 运行集成测试
go test -tags=integration ./...

# 生成测试覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 📦 部署

### Docker 部署

```bash
# 构建镜像
docker build -t codechange-tracker .

# 运行容器
docker run -d \
  --name codechange-tracker \
  -p 8080:8080 \
  -e DATABASE_URL="..." \
  codechange-tracker
```

### Kubernetes 部署

```bash
# 应用配置
kubectl apply -f deployments/k8s/
```

详细部署说明请参考：[Deployment Guide](docs/deployment.md)

## 🤝 贡献指南

我们欢迎社区贡献！请遵循以下步骤：

1. Fork 此仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开 Pull Request

### 代码贡献规范

- 遵循 Go 编码规范
- 编写单元测试
- 更新相关文档
- 提交信息遵循 [Conventional Commits](https://conventionalcommits.org/)

## 📄 许可证

本项目采用 MIT 许可证。详情请参考 [LICENSE](LICENSE) 文件。

## 🙏 致谢

感谢以下开源项目的支持：

- [Gin](https://github.com/gin-gonic/gin) - HTTP Web框架
- [GORM](https://github.com/go-gorm/gorm) - ORM库
- [PostgreSQL](https://www.postgresql.org/) - 数据库
- [Redis](https://redis.io/) - 缓存
- [Docker](https://www.docker.com/) - 容器化

## 📞 联系我们

- 项目主页: [GitHub Repository](https://github.com/xianshujun/CodeChangeTracker)
- 问题反馈: [GitHub Issues](https://github.com/xianshujun/CodeChangeTracker/issues)
- 文档: [Project Documentation](docs/)

---

<div align="center">
  <p>🌟 如果这个项目对你有帮助，请给我们一个Star！ 🌟</p>
</div>
