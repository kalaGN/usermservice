FROM golang:1.18.1-alpine

# 环境
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.io"

# 代码目录
WORKDIR /usermservice/

ADD . /usermservice/

WORKDIR /usermservice/server/v1

# 编译二进制文件
RUN go build -o service service.go 

# 声明服务端口
EXPOSE 8322

# 启动容器时运行的命令
CMD ["/usermservice/server/v1/service"]

