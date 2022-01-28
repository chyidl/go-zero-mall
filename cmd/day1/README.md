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
├── jarger                          # Jaeger 链路追踪
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