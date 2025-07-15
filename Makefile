# Makefile for CodeChangeTracker

.PHONY: help build run test clean docker-build docker-run setup-dev

# 默认目标
help:
	@echo "CodeChangeTracker Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  build          构建应用程序"
	@echo "  run            运行应用程序"
	@echo "  test           运行测试"
	@echo "  test-coverage  运行测试并生成覆盖率报告"
	@echo "  clean          清理构建文件"
	@echo "  setup-dev      设置开发环境"
	@echo "  docker-build   构建Docker镜像"
	@echo "  docker-run     运行Docker容器"
	@echo "  docker-compose 使用Docker Compose启动服务"
	@echo "  lint           代码检查"
	@echo "  fmt            格式化代码"
	@echo "  mod            整理模块依赖"

# 构建应用
build:
	@echo "构建应用程序..."
	go build -o bin/server cmd/server/main.go

# 运行应用
run:
	@echo "启动应用程序..."
	go run cmd/server/main.go

# 运行测试
test:
	@echo "运行测试..."
	go test -v ./...

# 运行测试并生成覆盖率报告
test-coverage:
	@echo "运行测试并生成覆盖率报告..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

# 清理构建文件
clean:
	@echo "清理构建文件..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# 设置开发环境
setup-dev:
	@echo "设置开发环境..."
	@if [ ! -f .env ]; then cp .env.example .env; echo "已创建.env文件，请编辑配置"; fi
	go mod download
	@echo "开发环境设置完成"

# 构建Docker镜像
docker-build:
	@echo "构建Docker镜像..."
	docker build -t codechange-tracker:latest .

# 运行Docker容器
docker-run: docker-build
	@echo "启动Docker容器..."
	docker run -d \
		--name codechange-tracker \
		-p 8080:8080 \
		--env-file .env \
		codechange-tracker:latest

# 使用Docker Compose启动服务
docker-compose:
	@echo "使用Docker Compose启动服务..."
	docker-compose up -d

# 停止Docker Compose服务
docker-compose-down:
	@echo "停止Docker Compose服务..."
	docker-compose down

# 代码检查
lint:
	@echo "运行代码检查..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint 未安装，使用基本检查..."; \
		go vet ./...; \
		if command -v golint >/dev/null 2>&1; then \
			golint ./...; \
		fi; \
	fi

# 格式化代码
fmt:
	@echo "格式化代码..."
	go fmt ./...

# 整理模块依赖
mod:
	@echo "整理模块依赖..."
	go mod tidy
	go mod verify

# 生成API文档
docs:
	@echo "生成API文档..."
	@if command -v swag >/dev/null 2>&1; then \
		swag init -g cmd/server/main.go; \
	else \
		echo "swag 未安装，请运行: go install github.com/swaggo/swag/cmd/swag@latest"; \
	fi

# 数据库迁移
migrate-up:
	@echo "执行数据库迁移..."
	@if command -v migrate >/dev/null 2>&1; then \
		migrate -path migrations -database "${DATABASE_URL}" up; \
	else \
		echo "migrate 工具未安装"; \
	fi

migrate-down:
	@echo "回滚数据库迁移..."
	@if command -v migrate >/dev/null 2>&1; then \
		migrate -path migrations -database "${DATABASE_URL}" down; \
	else \
		echo "migrate 工具未安装"; \
	fi

# 运行性能测试
bench:
	@echo "运行性能测试..."
	go test -bench=. -benchmem ./...

# 安全检查
security:
	@echo "运行安全检查..."
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "gosec 未安装，请运行: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"; \
	fi

# 开发工具安装
install-tools:
	@echo "安装开发工具..."
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# 完整的CI检查
ci: fmt lint test security
	@echo "所有CI检查完成"

# 本地完整构建
local-build: clean mod fmt lint test build
	@echo "本地构建完成"