# Rancher

[Rancher](https://docs.rancher.cn/) 是一套容器管理平台, 它可以帮助组织在生产环境中
轻松快捷的部署和管理容器. 也可以轻松地管理各种环境的Kubernetes.


## 安装
建议单独创建一个`kubernetes集群`(非托管集群), 该集群上只部署Rancher server. 安装完Rancher后, 
我们可使用以下两种方式引入集群, 用于部署自己的应用
- 使用Rancher创建新kubernetes集群
- 导入已有集群

### 测试使用(非生产)
**未实践**
- 安装
```
docker run -d --privileged --restart=unless-stopped \
  -p 80:80 -p 443:443 \
  rancher/rancher:latest
```