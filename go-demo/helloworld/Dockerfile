# 第一阶段：用于编译
# 指定用于编译的基础镜像
FROM golang:1.17-alpine AS builder

# 代理设置
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 指定当前工作目录（容器中项目存放目录）
WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod tidy

# 将当前目录下的所有文件拷贝至工作区 .
COPY . .

# 编译
RUN go build -o ./helloworld .

# 第二阶段：用于运行
# 指定用于运行的基础镜像
FROM alpine

WORKDIR /

COPY --from=builder /build/helloworld /

EXPOSE 8080

CMD ["/helloworld"]