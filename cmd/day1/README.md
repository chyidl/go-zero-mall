# Day 1
> 基于docker的go-zero运行环境

## 使用
> Docker本地开发环境搭建，使用Docker Compose编排管理我们的容器
```
go-zero-mall/cmd/day1 on  master [?] 
➜ tree .   
.
├── README.md
├── docker-compose.yaml
├── dtm                             # DTM 分布式事务管理器
│   ├── Dockerfile
│   └── config.yaml                 # DTM 配置文件
├── etcd                            # Etcd 服务注册发现
│   └── Dockerfile
├── golang                          # Golang 运行环境
│   └── Dockerfile
├── grafana                         # Grafana 可视化数据监控
│   └── grafana
├── jaeger                          # Jaeger 链路追踪
│   └── Dockerfile
├── mysql                           # MySQL 服务
│   └── Dockerfile
├── mysql-manage                    # MySQL 可视化管理
│   └── Dockerfile
├── prometheus                      # Prometheus 服务监控
│   ├── Dockerfile
│   └── prometheus.yml              # Prometheus 配置文件
├── redis                           # Redis 服务
│   └── Dockerfile
└── redis-manage                    # Redis 可视化管理
    └── Dockerfile

10 directories, 14 files
```

* 启动服务
```
docker-compose up -d
```

* 按需启动部分服务
```
docker-compose up -d etcd golang mysql redis
```

* 运行代码
> 将项目代码放置`CODE_PATH_HOST`制定的本机目录 进入

* 容器说明
| 容器名称 | 暴露端口 | host地址 | 说明 |
|:--------|:-------|:---------|:----|
| golang | 8000:8000<br>8001:8001<br>8002:8002<br>8003:8003<br>9000:9000<br>9001:9001<br>9002:9002<br>9003:9003<br>| golang | 80:开头的端口号用于api服务<br>90:开头的端口号用于rpc服务|
| dtm | 36789:36789<br>36790:36790 | dtm | dtm http协议和grpc协议服务端口号|
| etcd | 2379:2379 | etcd | etcd http api服务端口号|
| mysql | 3306:3306 | mysql | mysql 服务默认端口号|
| redis | 6379:6379 | redis | redis 服务默认端口号|
| prometheus | 3000:9090| prometheus | Prometheus web服务端口号|
| grafana | 4000:3000 | grafana | Grafana web 服务端口号 |
| jaeger | 5000:16686 | jaeger | Jaeger web 服务端口号|