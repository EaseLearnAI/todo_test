# 部署指南

## 📦 项目已准备就绪

项目已经完成构建并在本地运行成功！

### ✅ 本地测试

服务器正在运行：
- **URL**: http://localhost:5001
- **状态**: ✅ 运行中

你可以在浏览器中打开 http://localhost:5001 查看应用。

---

## 🚀 推送到 GitHub/GitLab

### 1️⃣ 创建远程仓库

在 GitHub 或 GitLab 上创建一个新仓库（例如：`todo-vue-flask`）

### 2️⃣ 添加远程仓库

```bash
cd /Users/terry/Desktop/coding/docker_learn/todo_test

# 添加 GitHub 远程仓库
git remote add origin https://github.com/你的用户名/todo-vue-flask.git

# 或添加 GitLab 远程仓库
git remote add origin https://gitlab.com/你的用户名/todo-vue-flask.git
```

### 3️⃣ 推送代码

```bash
git branch -M main
git push -u origin main
```

---

## 🌐 部署到 Vercel（推荐）

Vercel 原生支持 Python，但需要一些配置。

### 1️⃣ 创建 `vercel.json`

我已经为你准备好了 Vercel 配置文件（见下方）。

### 2️⃣ 安装 Vercel CLI（可选）

```bash
npm install -g vercel

# 登录
vercel login

# 部署
vercel
```

### 3️⃣ 或通过 Vercel 网站部署

1. 访问 https://vercel.com
2. 导入你的 Git 仓库
3. Vercel 会自动检测并部署

---

## 🐳 使用 Docker 部署

### 1️⃣ 构建镜像

```bash
docker build -t todo-vue-flask .
```

### 2️⃣ 运行容器

```bash
docker run -p 5001:5001 todo-vue-flask
```

---

## 📝 环境变量

部署时需要设置以下环境变量：

- `PORT`: 服务器端口（默认：5001）

---

## 🔧 本地开发命令

```bash
# 一键启动（推荐）
./start.sh

# 或手动启动

# 1. 激活虚拟环境
source venv/bin/activate

# 2. 构建前端
npm run build

# 3. 启动后端
python app.py
```

---

## 🎯 下一步操作

1. ✅ 项目已在本地成功运行
2. ⏭️ 创建远程 Git 仓库
3. ⏭️ 推送代码：`git push -u origin main`
4. ⏭️ 选择部署平台（Vercel/Heroku/Railway 等）

---

## 💡 提示

- 如果要修改端口，编辑 `.env` 文件中的 `PORT` 值
- 数据保存在 `data/todos.json` 文件中
- 前端构建产物在 `dist/` 目录
