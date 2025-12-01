# 前端构建阶段
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

# Go 后端构建阶段
FROM golang:1.21-alpine AS backend

WORKDIR /app

# 复制 Go 模块文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY main.go .

# 构建 Go 应用
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# 最终运行阶段
FROM alpine:latest

WORKDIR /app

# 安装 ca-certificates（用于 HTTPS）
RUN apk --no-cache add ca-certificates

# 从后端阶段复制编译好的二进制文件
COPY --from=backend /app/server .

# 从前端阶段复制构建产物
COPY --from=frontend /app/dist ./dist

# 复制环境变量文件
COPY .env .

# 创建数据目录
RUN mkdir -p data

# 暴露端口
EXPOSE 5001

# 设置环境变量
ENV PORT=5001

# 启动应用
CMD ["./server"]
