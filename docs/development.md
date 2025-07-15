# 开发环境设置指南

## 前置要求

- Go 1.21 或更高版本
- PostgreSQL 15+
- Redis 7+
- Docker 和 Docker Compose (可选)
- Git

## 快速开始

### 1. 克隆仓库

```bash
git clone https://github.com/xianshujun/CodeChangeTracker.git
cd CodeChangeTracker
```

### 2. 安装依赖

```bash
go mod download
```

### 3. 设置环境变量

复制环境变量示例文件：

```bash
cp .env.example .env
```

编辑 `.env` 文件，设置你的配置：

```env
DATABASE_URL=postgres://username:password@localhost:5432/codechange_tracker?sslmode=disable
REDIS_URL=redis://localhost:6379
JWT_SECRET=your-super-secret-key
OPENAI_API_KEY=your-openai-api-key
```

### 4. 启动数据库服务

#### 使用 Docker Compose (推荐)

```bash
# 启动数据库服务
docker-compose up -d db redis

# 等待服务启动
sleep 10
```

#### 手动安装数据库

如果你选择手动安装PostgreSQL和Redis，请确保它们正在运行并且配置正确。

### 5. 创建数据库

```bash
# 连接到PostgreSQL
psql -h localhost -U postgres

# 创建数据库
CREATE DATABASE codechange_tracker;
\q
```

### 6. 运行应用

```bash
go run cmd/server/main.go
```

应用将在 `http://localhost:8080` 启动。

### 7. 验证安装

访问健康检查端点：

```bash
curl http://localhost:8080/health
```

访问API文档：

```
http://localhost:8080/swagger/index.html
```

## 开发工作流

### 1. 项目结构

```
CodeChangeTracker/
├── cmd/                    # 应用入口点
│   └── server/
│       └── main.go
├── internal/               # 私有应用代码
│   ├── api/               # API层
│   │   ├── handlers/      # HTTP处理器
│   │   ├── middleware/    # 中间件
│   │   └── router.go      # 路由配置
│   ├── config/            # 配置管理
│   ├── database/          # 数据库相关
│   ├── models/            # 数据模型
│   ├── services/          # 业务逻辑
│   ├── analyzers/         # 代码分析器
│   └── utils/             # 工具函数
├── pkg/                   # 可重用的库代码
│   ├── logger/            # 日志包
│   ├── validator/         # 验证器
│   └── errors/            # 错误处理
├── docs/                  # 文档
├── test/                  # 测试文件
├── scripts/               # 脚本文件
├── deployments/           # 部署配置
├── docker-compose.yml     # Docker Compose配置
├── Dockerfile            # Docker镜像构建
├── go.mod               # Go模块定义
└── README.md            # 项目说明
```

### 2. 编码规范

- 遵循Go官方编码规范
- 使用 `gofmt` 格式化代码
- 使用 `golint` 检查代码质量
- 编写单元测试，覆盖率目标 >80%
- 为公共函数和方法编写文档注释

### 3. 提交规范

使用 Conventional Commits 规范：

```
feat: 添加新功能
fix: 修复bug
docs: 更新文档
style: 代码格式调整
refactor: 重构代码
test: 添加测试
chore: 构建过程或辅助工具的变动
```

### 4. 运行测试

```bash
# 运行所有测试
go test ./...

# 运行测试并生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### 5. 构建应用

```bash
# 构建应用
go build -o bin/server cmd/server/main.go

# 交叉编译
GOOS=linux GOARCH=amd64 go build -o bin/server-linux cmd/server/main.go
```

## 数据库迁移

### 手动迁移

应用启动时会自动执行数据库迁移。如果需要手动迁移：

```bash
go run cmd/migrate/main.go up
```

### 创建新迁移

```bash
# 安装migrate工具
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# 创建新迁移文件
migrate create -ext sql -dir migrations -seq add_new_table
```

## 调试

### 使用Delve调试器

```bash
# 安装Delve
go install github.com/go-delve/delve/cmd/dlv@latest

# 启动调试会话
dlv debug cmd/server/main.go
```

### IDE配置

#### VS Code

安装Go扩展，创建 `.vscode/launch.json`：

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Server",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/server/main.go",
            "env": {
                "ENVIRONMENT": "development"
            }
        }
    ]
}
```

#### GoLand

直接运行 `cmd/server/main.go` 文件。

## 常见问题

### 1. 数据库连接失败

确保PostgreSQL正在运行并且连接参数正确：

```bash
# 检查PostgreSQL状态
pg_isready -h localhost -p 5432

# 测试连接
psql -h localhost -U postgres -d codechange_tracker
```

### 2. Redis连接失败

检查Redis服务状态：

```bash
# 检查Redis状态
redis-cli ping
```

### 3. 端口被占用

更改 `.env` 文件中的端口设置：

```env
PORT=8081
```

### 4. Go模块问题

清理模块缓存：

```bash
go clean -modcache
go mod download
```

## 生产部署

### Docker部署

```bash
# 构建镜像
docker build -t codechange-tracker .

# 运行容器
docker run -d \
  --name codechange-tracker \
  -p 8080:8080 \
  -e DATABASE_URL="postgres://..." \
  -e REDIS_URL="redis://..." \
  codechange-tracker
```

### 使用Docker Compose

```bash
# 启动所有服务
docker-compose up -d

# 查看日志
docker-compose logs -f app

# 停止服务
docker-compose down
```

### Kubernetes部署

参考 `deployments/` 目录中的Kubernetes配置文件。

## 性能优化

### 1. 数据库优化

- 为频繁查询的字段添加索引
- 使用连接池
- 配置适当的超时时间

### 2. 缓存策略

- 使用Redis缓存频繁访问的数据
- 实施缓存过期策略
- 缓存API响应

### 3. 并发处理

- 使用goroutine处理耗时操作
- 实施任务队列
- 设置适当的并发限制

## 监控和日志

### 1. 应用监控

- 使用Prometheus收集指标
- 配置Grafana仪表盘
- 设置告警规则

### 2. 日志管理

- 使用结构化日志
- 集中化日志收集
- 日志级别配置

### 3. 健康检查

应用提供健康检查端点：

- `/health` - 基本健康检查
- `/health/live` - 存活性检查
- `/health/ready` - 就绪性检查