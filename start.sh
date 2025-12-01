#!/bin/bash

echo "🚀 启动 Todo 应用 (Go + Vue)..."

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo "❌ Go 未安装，请先安装 Go: https://golang.org/dl/"
    exit 1
fi

# 检查前端依赖
if [ ! -d "node_modules" ]; then
    echo "📦 安装前端依赖..."
    npm install
fi

# 构建前端
echo "🔨 构建前端..."
npm run build

# 下载 Go 依赖
echo "📦 下载 Go 依赖..."
go mod download

# 构建 Go 应用
echo "🔨 构建 Go 后端..."
go build -o server main.go

# 启动后端服务器
echo "✅ 启动服务器..."
echo "🌐 访问 http://localhost:5001"
./server
