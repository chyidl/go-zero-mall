# 服务监控

* Prometheus
> Prometheus 是一款基于时序数据库的开源监控告警系统，基本原理是通过HTTP协议周期性抓取被监控服务的状态，任意服务只要提供对应的HTTP接口就可以接入监控，不需要任何SDK或者其他集成过程，输出被监控服务信息的HTTP接口被叫做exporter.
    - 支持多维数据模型(由度量名和键值对组成的时间序列数据)
    - 不依赖分布式存储，单点服务器也可以使用
    - 支持HTTP协议主动拉取方式采集时间序列数据
    - 支持PushGateway推送时间序列数据
    - 支持服务发现和静态配置两种方式获取监控目标
    - 支持接入Grafana

* go-zero使用Prometheus监控服务

| 服务 | api 服务端口号 | rpc服务端口号 | api指标采集端口号 | rpc指标采集端口号 |
|:----|:--------------|:------------|:----------------|:---------------|
| user | 8000 | 9000| 9080| 9090|
| product | 8001| 9001| 9081| 9091|
| order | 8002| 9002| 9082| 9092 |
| pay | 8003| 9003| 9083| 9093|
