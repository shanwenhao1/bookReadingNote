# k8s 集群构建

## 初始化第一个master节点
- 初始化master节点, 使用[kubeadm-config.yaml](kubeadm-config.yaml), [参考](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init/#config-file) 
(可用命令`kubeadm config images list`查看该版本的镜像版本)
```bash
# 拉取镜像
kubeadm config images pull --config kubeadm-config.yaml
```