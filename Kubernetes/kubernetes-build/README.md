# Kubernetes集群搭建

k8s 版本为[v1.19](https://kubernetes.io/docs/home/)

## Master集群搭建
- 安装kubeadm、kubelet、kubectl
    ```bash
    sudo apt-get update && sudo apt-get install -y apt-transport-https curl
    ```
***
[apt-key.gpg](apt-key.gpg), ubuntu下使用[vpn client](vpn-use.md)
```bash
# google 官方源
# 如果出现gpg: no valid OpenPGP data found错误. 是因为需要翻墙下载apt-key.gpg. 
# 可选择访问https://packages.cloud.google.com/apt/doc/apt-key.gpg手动下载gpg文件. 地址: doc/kubernetes/files下
# 再执行apt-key add apt-key.gpg
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
```
***
    ```bash
    cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
    deb https://apt.kubernetes.io/ kubernetes-xenial main
    EOF
    # 阿里源
    echo "deb https://mirrors.aliyun.com/kubernetes/apt kubernetes-xenial main" >> /etc/apt/sources.list
    # 安装kubeadm等
    apt-get update
    # google 所用
    # apt-get install -y kubelet kubeadm kubectl
    apt-get install -y kubelet kubeadm kubectl --allow-unauthenticated
    # hold kubelet kubeadm kubectl的版本号, 不更新, 可使用unhold取消
    apt-mark hold kubelet kubeadm kubectl
    # 重启kubelet
    systemctl daemon-reload
    systemctl restart kubelet
    systemctl status kubelet
    ```

## 参考

- [k8s集群部署操作手册](k8s集群部署操作手册.pdf)
- [Ubuntu物理节点上部署kubernetes集群](https://www.kubernetes.org.cn/doc-17)
- [Kubeadm配置多Master](https://my.oschina.net/baobao/blog/3031712)