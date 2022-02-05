# 1 结构说明
## 1.1 go-demo

```sh
go-demo # 项目根目录
    ├── docker-compose.yml  # compose 配置文件，用于容器编排
    └── helloworld          # 项目目录
        ├── log             # 项目日志文件存放目录
        ├── Dockerfile      # docker 配置文件，用于镜像构建
        ├── .dockerignore   # docker 配置文件，用于构建时忽略不必要文件
        ├── go.mod
        └── main.go
```


## 1.2 Java-Tomcat-demo

```sh
Java-Tomcat-demo # 项目根目录
    ├── docker-compose.yml  # compose 配置文件，用于容器编排
    ├── question.md         # 部署过程中遇到的问题和解决方法
    └── helloworld          # 项目目录
        ├── log             # 项目日志文件存放目录
        ├── Dockerfile      # docker 配置文件，用于镜像构建
        └── webapps         # 存放war包的目录
```


## 1.3 Java-SpringBoot-demo

```sh
Java-SpringBoot-demo # 项目根目录
    ├── docker-compose.yml  # compose 配置文件，用于容器编排
    └── helloworld          # 项目目录
        ├── log             # 项目日志文件存放目录
        ├── Dockerfile      # docker 配置文件，用于镜像构建
        └── helloworld.jar  
```


# 2 运行说明
切换到相应项目根目录
> cd xxx-demo/

在 docker-compose.yml 同级目录下使用 docker-compose 构建
> docker-compose build

运行
> docker-compose up

后台运行
> docker-compose up -d
