# 服务拆分
> 一个商城可拆分用户服务(user),订单服务(order),产品服务(product),支付服务(pay),售后服务(afterSale)

* 用户服务(user)

| api 服务 | 端口: 8000 | rpc 服务 | 端口: 9000 |
|:--------|:------------|:----|:---|
| login | 用户登陆接口 | login | 用户登陆接口 |
| register | 用户注册接口 | register | 用户注册接口 |
| userinfo | 用户信息接口 | userinfo | 用户信息接口 |
| ... | ... | ... | ... |

* 产品服务(product)

| api 服务 | 端口: 8001 | rpc 服务 | 端口: 9001 |
|:---------|:-----------|:--------|:----------|
| create | 产品创建接口 | create | 产品创建接口 | 
| update | 产品修改接口 | update | 产品修改接口 |
| remove | 产品删除接口 | remove | 产品删除接口 | 
| detail | 产品详情接口 | detail | 产品详情接口 | 
| ... | ... | ... | ... |

* 订单服务 (order)

| api 服务 | 端口: 8002 | rpc服务 | 端口: 9002 |
|:---------|:----------|:--------|:----------|
| create | 订单创建接口 | create | 订单创建接口 ｜
| update | 订单修改接口 | update | 订单修改接口| 
| remove | 订单删除接口 | remove | 订单删除接口 | 
| detail | 订单详情接口 | detail | 订单详情接口 | 
| list | 订单列表接口 | list | 订单列表接口 |
| | | paid | 订单支付接口|
| ... | ... | ... | ... |

* 支付服务(pay)

| api 服务| 端口: 8003| rpc服务 | 端口: 9003 | 
|:-------|:----------|:-------|:----------|
| create | 支付创建接口| create | 支付创建接口|
| detail | 支付详情接口| detail | 支付详情接口|
| callback |支付回调接口| callback | 支付毁掉接口|
| ... | ... | ... | ... |


## 项目组织结构
```
go-zero-mall/cmd/mall on  master [!?] 
➜ tree .
.
├── common                  # 通用库
└── service                 # 服务
    ├── order
    │   ├── api             # order api 服务
    │   ├── model           # order 数据模型
    │   └── rpc             # order rpc 服务
    ├── pay
    │   ├── api             # pay api 服务
    │   ├── model           # pay 数据模型
    │   └── rpc             # pay rpc 服务
    ├── product
    │   ├── api             # product api 服务
    │   ├── model           # product 数据模型
    │   └── rpc             # product rpc 服务
    └── user
        ├── api             # user api 服务
        ├── model           # user 数据模型
        └── rpc             # user rpc 服务

18 directories, 0 files
```

## 服务拆分原则
- 由粗到细，避免过度拆分，遵循渐进式演进的原则
- 不同服务之间应该是正交的，不要你中有我我中有你
- 避免环形依赖，服务依赖关系应该是有向无环图
- 避免不同服务之间共享同一个数据库