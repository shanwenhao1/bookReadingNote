# Dashboard

[dashboard github](https://github.com/kubernetes/dashboard)

## 证书配置
注意: 由于Dashboard只允许本地使用HTTP连接进行访问, 其它地址只允许使用HTTPS访问. 因此这里kubectl proxy暴露外网访问需要 配合HTTPS方式访问. 
而外网访问则涉及到证书等问题.参照官方文档

## 插件安装
- 访问控制, [参考](https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/README.md)
- [helm 安装](https://artifacthub.io/packages/helm/k8s-dashboard/kubernetes-dashboard)
```bash
# Add kubernetes-dashboard repository
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
# Deploy a Helm Release named "my-release" using the kubernetes-dashboard chart
helm install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard
```
- 卸载
```
helm delete kubernetes-dashboard
```

## 参考
- [通过helm安装dashboard详细教程](https://www.cnblogs.com/baoshu/p/13326480.html#head2)