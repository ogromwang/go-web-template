# 说明
该项目为前些时间做的简单记账软件，项目结构简单，可作为 web 服务模板使用


# 技术选型
## 数据库
1. postgresql

## 对象存储
1. minio

## 鉴权
1. jwt

# 项目结构

```
.
├── Dockerfile
├── LICENSE
├── README.md
├── build.sh
├── cmd
│   └── main.go
├── config
│   ├── application.toml
│   └── db.toml
├── db
│   └── init.sql
├── go.mod
└── internal
    ├── config
    ├── dao
    ├── model
    ├── server
    ├── service
    └── util
```
