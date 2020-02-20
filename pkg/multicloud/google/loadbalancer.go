package google


/*
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