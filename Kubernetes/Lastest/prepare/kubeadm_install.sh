#!/bin/bash
# 更新并安装kubernetes `apt`所需包
sudo apt-get update
sudo apt-get install -y ca-certificates curl software-properties-common apt-transport-https curl ebtables ethtool iptables
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -
# 添加阿里源
echo "deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main" | sudo tee -a /etc/apt/sources.list.d/kubernetes.list
#对安装包进行签名
sudo apt-get update
# 安装kubeadm等
apt-get install -y kubelet kubeadm kubectl --allow-unauthenticated
# hold kubelet kubeadm kubectl的版本号, 不更新, 可使用unhold取消
apt-mark hold kubelet kubeadm kubectl
# 重启kubelet
systemctl daemon-reload
systemctl restart kubelet
