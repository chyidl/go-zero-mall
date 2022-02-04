# RPC服务Auth验证
* api
    - api服务中使用go-zero框架自带的jwt实现鉴权验证

* rpc
    - go-zero框架rpc服务的auth验证原理是客户端访问rpc服务需要携带App标识以及Token值，rpc服务会从指定的Redis服务中验证App标识和Token值是否正确。所以客户端的App标识,Token值，是需要提前打入Redis服务中.
