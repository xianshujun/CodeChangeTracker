#!/bin/bash

# CodeChangeTracker 开发环境初始化脚本

set -e

echo "🚀 CodeChangeTracker 开发环境初始化"
echo "=================================="

# 检查Go版本
echo "📋 检查Go版本..."
if ! command -v go &> /dev/null; then
    echo "❌ Go 未安装，请安装 Go 1.21 或更高版本"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
REQUIRED_VERSION="1.21"

if ! printf '%s\n%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V -C; then
    echo "❌ Go 版本不满足要求 (需要 >= $REQUIRED_VERSION，当前 $GO_VERSION)"
    exit 1
fi

echo "✅ Go 版本: $GO_VERSION"

# 检查Docker
echo "📋 检查Docker..."
if command -v docker &> /dev/null; then
    echo "✅ Docker 已安装"
    if command -v docker-compose &> /dev/null; then
        echo "✅ Docker Compose 已安装"
    else
        echo "⚠️  Docker Compose 未安装，建议安装以便使用"
    fi
else
    echo "⚠️  Docker 未安装，建议安装以便使用"
fi

# 创建配置文件
echo "📋 设置配置文件..."
if [ ! -f .env ]; then
    cp .env.example .env
    echo "✅ 已创建 .env 文件"
    echo "⚠️  请编辑 .env 文件设置你的配置"
else
    echo "✅ .env 文件已存在"
fi

# 安装Go依赖
echo "📋 安装Go依赖..."
go mod download
echo "✅ Go依赖安装完成"

# 创建必要的目录
echo "📋 创建项目目录..."
mkdir -p {bin,logs,tmp,test/fixtures}
echo "✅ 项目目录创建完成"

# 安装开发工具
echo "📋 安装开发工具..."
echo "安装 swag (API文档生成)..."
go install github.com/swaggo/swag/cmd/swag@latest

echo "安装 golangci-lint (代码检查)..."
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

echo "安装 gosec (安全检查)..."
go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest

echo "安装 migrate (数据库迁移)..."
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

echo "✅ 开发工具安装完成"

# 检查数据库连接
echo "📋 检查数据库配置..."
if grep -q "DATABASE_URL=postgres://username:password@localhost" .env; then
    echo "⚠️  请更新 .env 文件中的数据库配置"
else
    echo "✅ 数据库配置已设置"
fi

# 生成API文档
echo "📋 生成API文档..."
if command -v swag &> /dev/null; then
    swag init -g cmd/server/main.go
    echo "✅ API文档生成完成"
else
    echo "⚠️  swag 未正确安装，跳过API文档生成"
fi

# 运行基本测试
echo "📋 运行基本测试..."
if go test ./... > /dev/null 2>&1; then
    echo "✅ 基本测试通过"
else
    echo "⚠️  某些测试失败，这可能是正常的（如果数据库未配置）"
fi

echo ""
echo "🎉 开发环境初始化完成！"
echo ""
echo "下一步："
echo "1. 编辑 .env 文件设置你的配置"
echo "2. 启动数据库服务：docker-compose up -d db redis"
echo "3. 运行应用：make run 或 go run cmd/server/main.go"
echo "4. 访问 http://localhost:8080/health 检查服务状态"
echo "5. 访问 http://localhost:8080/swagger/index.html 查看API文档"
echo ""
echo "常用命令："
echo "  make help          - 查看所有可用命令"
echo "  make run           - 运行应用"
echo "  make test          - 运行测试"
echo "  make docker-compose - 使用Docker启动所有服务"
echo ""
echo "快乐编码！ 🚀"