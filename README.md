# 时间 Todo 应用 - Vue + Go

这是一个基于 Vue 3 前端和 Go 后端的 Todo 应用，从原 Node.js 版本改写而来。

## 📸 项目截图

> 一个简洁优雅的时间管理 Todo 应用

## 🛠 技术栈

### 前端
- **Vue 3** - 渐进式 JavaScript 框架
- **Vite** - 下一代前端构建工具
- **原生 CSS** - 无需任何 UI 框架

### 后端
- **Go 1.21+** - 高性能编译型语言
- **Gorilla Mux** - 强大的 HTTP 路由器
- **CORS 中间件** - 跨域资源共享支持
- **JSON 文件存储** - 简单可靠的数据持久化

## ✨ 功能特性

- ✅ 创建、编辑、删除 Todo
- ✅ 设置截止时间（datetime-local 选择器）
- ✅ 切换完成状态（复选框）
- ✅ 自动排序（未完成优先，按时间排序）
- ✅ 实时保存到 JSON 文件
- ✅ 响应式设计，适配移动端

## 🚀 快速开始

### 方式一：一键启动（推荐）

```bash
./start.sh
```

### 方式二：手动启动

#### 1. 安装 Go

确保已安装 Go 1.21 或更高版本：
```bash
go version
```

如未安装，访问：https://golang.org/dl/

#### 2. 安装前端依赖

```bash
npm install
```

#### 3. 构建前端

```bash
npm run build
```

#### 4. 下载 Go 依赖

```bash
go mod download
```

#### 5. 构建并运行 Go 服务器

```bash
go build -o server main.go
./server
```

#### 6. 访问应用

打开浏览器访问：http://localhost:5001

## 💻 开发模式

如果你想同时开发前端和后端，可以分别启动：

**终端 1 - 后端（Go）**:
```bash
go run main.go
# 后端运行在 http://localhost:5001
```

**终端 2 - 前端（Vite 热重载）**:
```bash
npm run dev
# 前端开发服务器运行在 http://localhost:3000
# API 请求会自动代理到后端
```

## 📡 API 接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/todos` | 获取所有 Todo |
| POST | `/api/todos` | 创建新 Todo |
| PUT | `/api/todos/:id` | 更新 Todo |
| POST | `/api/todos/:id/toggle` | 切换完成状态 |
| DELETE | `/api/todos/:id` | 删除 Todo |

### 请求示例

```bash
# 获取所有 Todos
curl http://localhost:5001/api/todos

# 创建新 Todo
curl -X POST http://localhost:5001/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"完成项目","dueAt":"2025-12-31T23:59"}'

# 切换完成状态
curl -X POST http://localhost:5001/api/todos/{id}/toggle

# 删除 Todo
curl -X DELETE http://localhost:5001/api/todos/{id}
```

## 📁 项目结构

```
todo_test/
├── src/                  # Vue 源代码
│   ├── App.vue          # 主组件
│   ├── main.js          # 入口文件
│   └── style.css        # 样式文件
├── dist/                # 前端构建产物（自动生成）
├── data/                # 数据存储目录
│   └── todos.json       # Todo 数据文件
├── main.go              # Go 后端主文件
├── go.mod               # Go 模块定义
├── go.sum               # Go 依赖锁定
├── package.json         # Node.js 依赖
├── vite.config.js       # Vite 配置
├── index.html           # HTML 模板
├── Dockerfile           # Docker 配置
├── vercel.json          # Vercel 部署配置
├── start.sh             # 一键启动脚本
├── deploy.sh            # 部署脚本
├── DEPLOY.md            # 详细部署指南
└── README.md            # 项目说明（本文件）
```

## 🐳 Docker 部署

### 构建镜像

```bash
docker build -t todo-vue-flask .
```

### 运行容器

```bash
docker run -p 5001:5001 todo-vue-flask
```

访问 http://localhost:5001

## 🌐 在线部署

### 部署到 Vercel

1. 推送代码到 GitHub
2. 在 Vercel 导入仓库
3. 自动部署完成

详细步骤请查看 [`DEPLOY.md`](./DEPLOY.md)

### 部署到其他平台

- **Heroku**: 支持 Python，配置 `Procfile`
- **Railway**: 自动检测 Python 项目
- **Render**: 支持 Docker 部署

## 🔧 配置说明

### 环境变量 (`.env`)

```env
PORT=5001  # 服务器端口
```

### Vite 代理配置 (`vite.config.js`)

开发模式下，前端 API 请求会自动代理到后端：

```javascript
proxy: {
  '/api': {
    target: 'http://localhost:5001',
    changeOrigin: true
  }
}
```

## 📝 开发笔记

### 从 Node.js 到 Vue + Go 的改写要点

1. **前端框架化**
   - 原：原生 JavaScript DOM 操作
   - 新：Vue 3 响应式数据绑定

2. **后端语言切换**
   - 原：Node.js + Express 风格
   - 新：Go + Gorilla Mux RESTful API

3. **性能优化**
   - Go 的编译型特性带来更高的执行效率
   - 内置并发支持，goroutine 轻量级线程
   - 更小的内存占用

4. **部署便利性**
   - 单一二进制文件，无需运行时环境
   - Docker 镜像更小（使用 Alpine 基础镜像）

## 🎯 Git 推送命令

### 初始化并推送到 GitHub

```bash
# 1. 创建 GitHub 仓库（在网页上操作）

# 2. 添加远程仓库
git remote add origin https://github.com/你的用户名/todo-vue-flask.git

# 3. 推送代码
git branch -M main
git push -u origin main
```

### 或使用快捷脚本

```bash
./deploy.sh
```

## 🐛 故障排除

### 端口被占用

```bash
# macOS 关闭 AirPlay Receiver（占用 5000 端口）
# 系统偏好设置 -> 通用 -> 隔空播放接收器 -> 关闭

# 或修改 .env 文件使用其他端口
PORT=5001
```

### Python 包安装失败（代理问题）

该问题已不存在，使用 Go 不需要 Python 环境。

### Go 依赖下载失败

```bash
# 设置 Go 代理（国内用户）
export GOPROXY=https://goproxy.cn,direct
go mod download
```

### 前端构建失败

```bash
# 清除缓存重新安装
rm -rf node_modules package-lock.json
npm install
npm run build
```

## 📄 许可证

MIT License

## 🙏 致谢

- 原项目基于 Node.js 实现
- 使用 Vue 3 和 Flask 重写
- 感谢开源社区的支持

---

**享受使用！如有问题欢迎提 Issue** 🎉
