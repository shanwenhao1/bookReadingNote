# k8s 集群构建

## 初始化第一个master节点
- 初始化master节点, 使用[kubeadm-config.yaml](kubeadm-config.yaml), [参考](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init/#config-file) 
(可用命令`kubeadm config images list`查看该版本的镜像版本)
    - 拉取镜像
        ```bash
        # 拉取镜像
        kubeadm config images pull --config kubeadm-config.yaml
        ```
        - 部分镜像拉取不到的问题, [参考](https://www.jianshu.com/p/78601ae3e988)
            ```bash
            # 部分镜像拉取不到, 自己手动拉取
            docker pull coredns/coredns
            docker tag coredns/coredns:latest registry.aliyuncs.com/google_containers/coredns/coredns:v1.8.0
            ```
    - 初始化master节点 
        ```bash
        # 初始化第一个集群节点
        kubeadm init --config kubeadm-config.yaml
        ```
      - 保存`kubeadm join`命令
          ```bash
          # 如果忘记token可用命令 'kubeadm token list'查看已有token, 或者新建新的token
          # kubeadm join 命令格式为:
          kubeadm join --token <token> <control-plane-host>:<control-plane-port> --discovery-token-ca-cert-hash sha256:<hash>
          # 其中--discovery-token-ca-cert-hash的值可以通过以下命令获得
          openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | \
             openssl dgst -sha256 -hex | sed 's/^.* //'
          ```
          - master:
              ```bash
              kubeadm join k8s.swh.com:6443 --token k8sxlq.2tziv5dd877y3an2 \
              	--discovery-token-ca-cert-hash sha256:6fb97faccc431fd6debd8d0b405e53ffd110f934cfa6f27a688991c84a9a6dec \
              	--control-plane 
              ```
          - node:
              ```
              kubeadm join k8s.swh.com:6443 --token k8sxlq.2tziv5dd877y3an2 \
              	--discovery-token-ca-cert-hash sha256:6fb97faccc431fd6debd8d0b405e53ffd110f934cfa6f27a688991c84a9a6dec
              ```
      ![](picture/master%20init.png)
    - 将kubeConfig置入环境变量
        ```bash
        # 临时
        export KUBECONFIG=/etc/kubernetes/admin.conf
        # 永久
        echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> /root/.bashrc
        source /root/.bashrc
        ```
    - 安装calico网络插件, [参考](https://docs.projectcalico.org/getting-started/kubernetes/quickstart)
        - 下载calico网络插件配置信息, 并修改`CALICO_IPV4POOL_CIDR`值为`kubeadm-config.yaml`的podSubnet值.
        修改后的[calico.yaml](../../kubernetes-build/yaml/calico.yaml)文件
        ```bash
        wget https://docs.projectcalico.org/manifests/calico.yaml
        # 修改后安装calico网络插件
        kubectl apply -f calico.yaml
        ```
    - 至此, 第一个master节点搭建完成
        ![](picture/first%20master%20okay.png)

## 其他节点加入集群
### master节点加入
- 使用[脚本](../../kubernetes-build/sh/sync.master.ca.sh)将ca证书从第一个master节点拷贝至其他master节点
- 使用`kubeadm join`命令加入master或者node节点
- 添加环境变量
```bash
echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> /root/.bashrc
source /root/.bashrc
```
### node节点加入
- 使用`kubeadm join`命令加入master或者node节点
- 如果想要在节点上使用`kubectl`命令, 将`/etc/kubernetes/admin.conf`复制到节点
```bash
scp /etc/kubernetes/admin.conf root@192.168.1.115:/etc/kubernetes/
# 将kubeConfig永久置入环境变量
echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> /root/.bashrc
source /root/.bashrc
```