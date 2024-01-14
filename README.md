### 介绍
请求调用过程: controller ---> logic ---> dao ---> 外部服务等

```
.
├── Makefile
├── README.md
├── config                  # 全局配置文件
│   ├── config.go
│   └── service.conf
├── controller              # 请求入口，负责初始化api、对请求做基本的校验
│   ├── router
│   │   ├── middleware
│   │   │   └── trace.go
│   │   └── router.go
│   └── user
│       ├── handle.go
│       ├── register.go
│       └── user.go
├── dao                     # 对数据库封装的增删改查等通用接口
│   ├── database.go
│   └── user.go
├── go.mod
├── go.sum
├── logic                   # 项目的核心逻辑处理接口
│   └── user.go
├── main.go
└── utils                   # 一些通用的、业务无关的小函数
    ├── fstools
```