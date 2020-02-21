package google

/*
// https://cloud.google.com/load-balancing/docs/choosing-load-balancer?hl=zh_CN
Cloud 负载平衡器摘要
下表提供了有关各负载平衡器的一些具体信息。

负载平衡器	流量类型	保留客户端 IP 地址	全球性或区域性	负载平衡方案	负载平衡器目标端口	代理或直通
HTTP(S)	HTTP 或 HTTPS	否	全球	EXTERNAL	在端口 80 或 8080 上处理 HTTP 流量；在端口 443 上处理 HTTPS 流量	代理
SSL 代理	具有 SSL 分流的 TCP	否	全球	EXTERNAL	25、43、110、143、195、443、465、587、700、993、995、1883、5222	代理
TCP 代理	无 SSL 分流的 TCP	否	全球	EXTERNAL	25、43、110、143、195、443、465、587、700、993、995、1883、5222	代理
网络 TCP/UDP	TCP 或 UDP	是	区域	EXTERNAL	不限	直通
内部 TCP/UDP	TCP 或 UDP	是	区域	INTERNAL	不限	直通
内部 HTTP(S)	HTTP 或 HTTPS	否	区域	INTERNAL_MANAGED	在端口 80 或 8080 上处理 HTTP 流量；在端口 443 上处理 HTTPS 流量	代理


// loadbalancer 特性
https://cloud.google.com/load-balancing/docs/features?hl=zh_CN

loadbalancer 相关的接口：
负载均衡列表：
https://stackoverflow.com/questions/50814300/gcp-api-for-getting-list-of-load-balancer


转发规则【转发策略】：
【1】内网HTTP/HTTPS负载均衡 https://cloud.google.com/compute/docs/reference/rest/v1/urlMaps/insert?hl=zh_CN
【2】外网HTTP/HTTPS负载均衡 https://cloud.google.com/compute/docs/reference/rest/v1/regionUrlMaps/insert?hl=zh_CN


健康检查：
https://cloud.google.com/compute/docs/reference/rest/v1/healthChecks?hl=zh_CN
{
  "checkIntervalSec": 10.0,
  "creationTimestamp": "2020-02-17T23:58:35.894-08:00",
  "description": "",
  "healthyThreshold": 2.0,
  "id": "3750627300595869636",
  "kind": "compute#healthCheck",
  "name": "tb-elb-http-healthcheck",
  "region": "projects/qiujian-yunion-hk/regions/asia-east1",
  "selfLink": "projects/qiujian-yunion-hk/regions/asia-east1/healthChecks/tb-elb-http-healthcheck",
  "tcpHealthCheck": {
    "port": 80.0,
    "proxyHeader": "NONE"
  },
  "timeoutSec": 5.0,
  "type": "TCP",
  "unhealthyThreshold": 3.0
}

[1] normal health check
https://cloud.google.com/compute/docs/reference/rest/v1/healthChecks/insert?hl=zh_CN
[2] http health check
https://cloud.google.com/compute/docs/reference/rest/v1/httpHealthChecks/insert?hl=zh_CN
[3] https health check
https://cloud.google.com/compute/docs/reference/rest/v1/httpsHealthChecks/insert?hl=zh_CN

*/
