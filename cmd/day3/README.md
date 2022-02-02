# 用户服务

## 生成 user model 模型
```
# 进入服务工作区
$ cd mall/service/user

# 创建 sql 文件
$ vim model/user.sql

# 模版生成命令
go-zero-mall/cmd/mall/service/user on  master [!?] 
➜ goctl model mysql ddl -src ./model/user.sql -dir ./model -c
Done.
```

## 生成 user api 服务
```
# 创建 api 文件
$ vim api/user.api

# 运行模版生成命令
$ goctl api go -api ./api/user.api -dir ./api
```

## 生成 user rpc 服务
```
# 创建 proto 文件
$ vim rpc/user.proto

# 运行模版生成命令
$ goctl rpc proto -src ./rpc/user.proto -dir ./rpc
```

## 编写配置文件
