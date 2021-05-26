# Helm
- [Helm详解](#Helm详解)
    - [Helm使用教程](#Helm使用教程)
    - [Helm组件](#Helm组件)
- [安装及使用](#安装及使用)
    - [安装](#安装)
        - [Apt安装](#Apt安装)
        - [Snap安装](#Snap安装)
        - [二进制安装](#二进制安装)
    - [Helm使用](#Helm使用)

## Helm详解
Helm是kubernetes的一个包管理工具, 它具有如下功能额:
- 创建新的chart
- chart打包成tgz格式
- 上传chart到chart仓库或从仓库中下载chart
- 在`kubernetes`集群中安装或卸载chart
- 管理用`Helm`安装的chart的发布周期

Helm有三个重要概念:
- chart: 包含了创建`kubernetes`的一个应用实例的必要信息
- config: 包含了应用发布配置信息
- release: 是一个chart及其配置的一个运行实例

### Helm使用教程
[官方文档](https://helm.sh/zh/docs/intro/quickstart/)

#### 仓库
Helm的Repo仓库和Docker Registry比较相似, Chart库可以用来存储和共享打包 Chart 的位置, 安装Helm后, 默认仓库地址是
google的一个地址, 无法科学上网的无法访问到官方提供的Chart仓库．使用`helm repo list`查看当前的仓库配置.

Chart 仓库其实就是一个带有`index.yaml`索引文件和任意个打包的Chart的HTTP服务器而已


### Helm组件
Helm 3移除了Tiller, 使用与kubectl上下文相同的访问权限

## 安装及使用

### 安装
[参考](https://helm.sh/zh/docs/intro/install/)

#### Apt安装
```bash
curl https://baltocdn.com/helm/signing.asc | sudo apt-key add -
sudo apt-get install apt-transport-https --yes
echo "deb https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list
sudo apt-get update
sudo apt-get install helm
```

#### Snap安装
```bash
sudo snap install helm --classic
```

#### 二进制安装
- 在[Helm Release](https://github.com/helm/helm/releases) 下载二进制文件, 解压后将执行文件`helm`拷贝到
`/usr/local/bin`目录下即可.
```bash
# 在安装Helm的机器上查看版本号
helm version
```
- 添加chart仓库, [kubernetes官方查找chart包的地址](https://hub.kubeapps.com/charts)
```bash
# helm 官方chart仓库
helm repo add stable https://charts.helm.sh/stable
# kubernetes的官方chart仓库(国内无法访问)
helm repo add stable https://kubernetes-charts.storage.googleapis.com/
# 阿里源https://github.com/cloudnativeapp/charts
helm repo add apphub https://apphub.aliyuncs.com
```
  
### Helm使用
- install相关
```bash
# Make sure we get the latest list of charts
helm repo update
helm install stable/mysql --generate-name
# 查看chart包信息
helm show chart stable/mysql
# 卸载
helm uninstall smiling-penguin
# 状态
helm status smiling-penguin
```