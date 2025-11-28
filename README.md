# 时间 Todo 应用 - Vue + Flask

这是一个基于 Vue 3 前端和 Flask 后端的 Todo 应用。

## 技术栈

- **前端**: Vue 3 + Vite
- **后端**: Flask + Flask-CORS
- **数据存储**: JSON 文件

## 功能特性

- ✅ 创建、编辑、删除 Todo
- ✅ 设置截止时间
- ✅ 切换完成状态
- ✅ 按时间和完成状态排序

## 本地开发

### 1. 安装后端依赖

```bash
# 创建 Python 虚拟环境（推荐）
python3 -m venv venv
source venv/bin/activate  # macOS/Linux
# 或 venv\Scripts\activate  # Windows

# 安装依赖
pip install -r requirements.txt
```

### 2. 安装前端依赖

```bash
npm install
```

### 3. 启动开发服务器

**终端 1 - 启动后端**:
```bash
python app.py
# 后端运行在 http://localhost:5000
```

**终端 2 - 启动前端**:
```bash
npm run dev
# 前端运行在 http://localhost:3000
```

### 4. 构建生产版本

```bash
# 构建前端
npm run build

# 启动生产服务器（Flask 会提供构建后的静态文件）
python app.py
```

## 部署到 Vercel

1. 初始化 Git 仓库并推送代码
2. 在 Vercel 导入项目
3. 配置构建命令和输出目录

## API 接口

- `GET /api/todos` - 获取所有 Todo
- `POST /api/todos` - 创建新 Todo
- `PUT /api/todos/:id` - 更新 Todo
- `POST /api/todos/:id/toggle` - 切换完成状态
- `DELETE /api/todos/:id` - 删除 Todo

## 项目结构

```
todo_test/
├── src/              # Vue 源代码
│   ├── App.vue      # 主组件
│   ├── main.js      # 入口文件
│   └── style.css    # 样式文件
├── app.py           # Flask 后端
├── requirements.txt # Python 依赖
├── package.json     # Node 依赖
├── vite.config.js   # Vite 配置
└── index.html       # HTML 模板
```
