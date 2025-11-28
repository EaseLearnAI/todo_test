#!/bin/bash

# 快速部署脚本

set -e

echo "🚀 开始部署流程..."
echo ""

# 1. 检查 Git 状态
echo "📋 检查 Git 状态..."
if ! git diff-index --quiet HEAD --; then
    echo "⚠️  有未提交的更改"
    git status --short
    read -p "是否提交这些更改? (y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        read -p "请输入提交信息: " commit_msg
        git add .
        git commit -m "$commit_msg"
    fi
fi

# 2. 检查远程仓库
echo ""
echo "🔗 检查远程仓库..."
if ! git remote get-url origin > /dev/null 2>&1; then
    echo "❌ 未设置远程仓库"
    echo ""
    echo "请先添加远程仓库："
    echo "  git remote add origin <你的仓库URL>"
    echo ""
    echo "示例："
    echo "  git remote add origin https://github.com/username/todo-vue-flask.git"
    exit 1
fi

REMOTE_URL=$(git remote get-url origin)
echo "✅ 远程仓库: $REMOTE_URL"

# 3. 推送代码
echo ""
echo "📤 推送代码到远程仓库..."
git push -u origin main

echo ""
echo "✅ 代码推送成功！"
echo ""
echo "🎉 下一步："
echo "  1. 访问你的 Git 仓库查看代码"
echo "  2. 在 Vercel/Heroku/Railway 等平台导入仓库进行部署"
echo "  3. 或使用 Docker 部署: docker build -t todo-app . && docker run -p 5001:5001 todo-app"
echo ""
