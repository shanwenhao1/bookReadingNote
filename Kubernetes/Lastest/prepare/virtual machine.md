# 虚拟机创建

## 虚拟机配置
### 三台`master`节点主机
- ubuntu20.04+
- 每台机器 2 GB 或更多的 RAM （如果少于这个数字将会影响你应用的运行内存)
- 2 CPU 核或更多
- 集群中的所有机器的网络彼此均能相互连接(公网和内网都可以)
- 节点之中不可以有重复的主机名、MAC 地址或 product_uuid. 以及[所需开放的端口](https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/install-kubeadm/#check-required-ports)
    ```bash
    # 检查MAC address唯一性
    ifconfig -a
    # 检查product_uuid唯一性
    sudo cat /sys/class/dmi/id/product_uuid
    ```
- 禁用交换分区。为了保证`kubelet`正常工作, 你必须禁用交换分区
    ```bash
    # 临时关闭
    swapoff -a
    # 或永久关闭
    vim /etc/fstab
    #/swap.img       none    swap    sw      0       0
    ```
- 允许iptables检查桥接流量
    ```bash
    cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
    br_netfilter
    EOF
    
    cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
    net.bridge.bridge-nf-call-ip6tables = 1
    net.bridge.bridge-nf-call-iptables = 1
    EOF
    
    sudo sysctl --system
    ```
  
### 三台+`node`节点主机

## 通用设置
- 虚拟机使用桥接模式
- 创建完成后, 允许root远程登录
```
# 更改设置允许root远程登录
echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
# 重启sshd服务
sudo systemctl restart sshd
# 修改root密码
sudo passwd root
```
- 添加阿里源
```bash
echo "deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse" >> /etc/apt/sources.list

apt-get update
```
- 添加必要的工具
```bash
apt-get -y install lrzsz
```
- 固定虚拟机ip, 使用静态ip, [参考](https://www.cnblogs.com/kehoudaanxianjie/p/13139636.html)
    - 修改`/etc/netplan`目录下的`.yaml`文件, 修改成静态ip模式
        ```bash
        # 更改配置
        echo "# This is the network config written by 'subiquity'
        network:
          ethernets:
            ens33:
              dhcp4: no
              addresses: [192.168.1.113/24]
              optional: true
              gateway4: 192.168.1.1
              nameservers:
                      addresses: [223.5.5.5,223.6.6.6]
          version: 2" > /etc/netplan/00-installer-config.yaml
        
        # 应用配置
        netplan apply
        ```
        ![](picture/netplan%20setting.png)
        -  设置的ip如下
            - `k8s master 1` ip为`192.168.1.112`
            - `k8s master 2` ip为`192.168.1.113`
            - `k8s master 3` ip为`192.168.1.114`
            - `k8s node 1` ip为`192.168.1.115.`
            - `k8s node 2` ip为`192.168.1.116`
            - `k8s node 3` ip为`192.168.1.117`
            
### Docker安装
[文档](../../prepare/docker.md)


## 安装kubeadm、kubectl、kubelet
- Google官方源(未采用)
```bash
# 更新并安装kubernetes `apt`所需包
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl
# 下载Google Cloud公开签名秘钥
sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
# 添加kubernetes apt仓库
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
# 安装
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```
- 阿里源, [参考](https://developer.aliyun.com/mirror/kubernetes), 
[kubeadm_install.sh](kubeadm_install.sh)
```bash
# 安装kubeadm、kubectl、kubelet
sh kubeadm_install.sh
systemctl status kubelet
```
- snap 安装(未采用)
```bash
# 参考, 未实践
snap install kubeadm --classic
```

## 模拟域名解析
- 将`/etc/hostname`内主机名修改成`k8s-master-1`(相应的主机改为其对应名称, k8s-master-1、k8s-master-2...)
```bash
# 也可用命令
hostnamectl set-hostname k8s-master-1
```
- 修改`/etc/hosts`
```bash
127.0.1.1 k8s-master-1

echo "192.168.1.112 k8s-master-1
192.168.1.113 k8s-master-2
192.168.1.114 k8s-master-3

192.168.1.115 k8s-node-1
192.168.1.116 k8s-node-2
192.168.1.117 k8s-node-3

# master节点域名解析涉及负载均衡(本示例并没有实现LB, 生产环境中至少三台) 

192.168.1.112  k8s.swh.com
192.168.1.113  k8s.swh.com
192.168.1.114  k8s.swh.com" >> /etc/hosts
```