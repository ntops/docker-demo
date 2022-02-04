# go-demo
## 结构说明
```sh
go-demo # 项目根目录
    ├── docker-compose.yml  # compose 配置文件，用于容器编排
    └── helloworld          # 项目目录
        ├── log             # 项目日志文件存放目录
        ├── Dockerfile      # docker 配置文件，用于镜像构建
        ├── .dockerignore   # docker 配置文件，用于构建时忽略不必要文件
        ├── go.mod
        ├── go.sum
        └── main.go
```
## 运行说明
切换到项目根目录
> cd go-demo/

使用 docker-compose 构建
> docker-compose build

运行
> docker-compose up

后台运行
> docker-compose up -d
