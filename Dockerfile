# 使用官方Golang镜像作为基础镜像
FROM golang:latest
# 设置工作目录
WORKDIR /app
# 将当前目录下的所有文件复制到容器的/app目录下
ADD . /app
# 编译Golang代码
RUN go build -o time-manager
# 设置容器启动命令
CMD ["sh", "-c", "./time-manager"]