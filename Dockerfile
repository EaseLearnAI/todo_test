FROM node:18-alpine AS frontend

WORKDIR /app

# 复制前端文件
COPY package*.json ./
COPY vite.config.js ./
COPY index.html ./
COPY src ./src

# 安装依赖并构建
RUN npm install
RUN npm run build

# Python 后端阶段
FROM python:3.9-slim

WORKDIR /app

# 安装 Python 依赖
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# 复制后端代码
COPY app.py .
COPY .env .

# 从前端阶段复制构建产物
COPY --from=frontend /app/dist ./dist

# 创建数据目录
RUN mkdir -p data

# 暴露端口
EXPOSE 5001

# 设置环境变量
ENV PORT=5001

# 启动应用
CMD ["python", "app.py"]
