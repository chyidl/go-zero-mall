# Jaeger
> 开源的分布式追踪系统,兼容OpenTracing API.

    - 分布式追踪信息传递
    - 分布式事务监控
    - 问题分析
    - 服务依赖性分析
    - 性能优化

* Jaeger 的全链路追踪功能主要由三个角色完成
    - client: 负责全链路上各个调用点的计时、采样、并将tracing数据法网本地的agent.
    - agent: 负责收集 client 发来的 tracing 数据，并以thrift协议转发给collector.
    - collector: 负责搜集所有agent上报的tracing数据，统一存储