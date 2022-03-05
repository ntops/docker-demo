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

## 1.4 wordpress-demo

```sh
wordpress-demo # 项目根目录
    ├── docker-compose.yml  # compose 配置文件，用于容器编排
    └── helloworld          # 项目目录
        ├── docker-entrypoint.sh、Dockerfile、wp-config-docker.php # 在自己构建镜像时使用
        └── wp-content     # wordpress 存放数据的目录
```


# 2 运行说明
切换到相应项目根目录
> cd xxx-demo/

在 docker-compose.yml 同级目录下使用 docker-compose 构建
> docker compose build

运行
> docker compose up

后台运行
> docker compose up -d

重启
> docker compose restart

停止并删除
> docker compose down

注：其他命令自行百度

# 3 demo 扩展说明
1、**关于端口**：demo 运行时监听本地 8080 端口并映射到容器内部项目相关端口

2、**关于数据库**：若有 mysql 需求可联系管理员分配相关权限，项目可通过本地映射到容器的ip 172.17.0.1 连接到本地 mysql 服务

3、待续...
