#!/bin/bash

echo "🚀 启动 Todo 应用..."

# 检查虚拟环境
if [ ! -d "venv" ]; then
    echo "📦 创建 Python 虚拟环境..."
    python3 -m venv venv
fi

# 激活虚拟环境
source venv/bin/activate

# 安装后端依赖
echo "📦 安装后端依赖..."
pip install -r requirements.txt

# 检查前端依赖
if [ ! -d "node_modules" ]; then
    echo "📦 安装前端依赖..."
    npm install
fi

# 构建前端
echo "🔨 构建前端..."
npm run build

# 启动后端服务器
echo "✅ 启动服务器..."
echo "🌐 访问 http://localhost:5000"
python app.py
