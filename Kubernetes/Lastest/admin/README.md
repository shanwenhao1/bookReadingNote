# 集群管理配置

## 证书(不使用https的话了解就行)

Kubernetes集群中使用HTTPS协议, 需要以下几个部件
- 证书管理器: 这里使用 [cert-manager](cert-manager/README.md) (原生的kubernetes证书管理器)
- 一个证书自动签发服务: 通过ingress来发布HTTPS服务
    - 依赖: Ingress Controller进行配置, 用来启用HTTPS及其路由


### 参考
[kubernets 手动生成证书官方教程](https://kubernetes.io/zh/docs/tasks/administer-cluster/certificates/#openssl)


## ingress控制器
[安装](ingress/README.md)

## 集群管理web界面

### Dashboard
`Kubernetes Dashboard` 是基于网页的Kubernetes用户界面, 用于管理和浏览集群资源和应用信息.
[官方文档](https://kubernetes.io/zh/docs/tasks/access-application-cluster/web-ui-dashboard/)

[安装教程](dashboard/README.md)

### prometheus监控报警系统
`Prometheus` 是Google内部监控报警系统的开源版本, 是Google SRE思想不断完善的产物. 
它是用来替代kubernetes早期的`heapster`、`influxDB`、`grafana`组合来监控系统.

[文档](prometheus/README.md)

**TODO继续**