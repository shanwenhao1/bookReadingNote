# cert-manager

[cert-manager](https://cert-manager.io/docs/) 是kubernetes原生的证书管理器
![](picture/high-level-overview.svg)


## 安装
[helm 安装官方教程](https://cert-manager.io/docs/installation/kubernetes/)
- cert-manager 要求一系列的CRD resources
    - 方式1: 使用kubectl安装CRDs
        ```bash
        kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.3.1/cert-manager.crds.yaml
        ```
    - 方式2(采用): helm安装时使用`--set installCRDs=true`, [参数详情](https://artifacthub.io/packages/helm/cert-manager/cert-manager)
- helm 安装cert-manager
    ```bash
    kubectl create namespace cert-manager
    helm repo add jetstack https://charts.jetstack.io
    helm repo update
    helm install cert-manager jetstack/cert-manager --namespace cert-manager --version v1.3.1 --set installCRDs=true
    ```
    - 查看部署情况
        ```bash
        kubectl get pods --namespace cert-manager
        ```
    - 部署完成后, 测试 webhook 工作是否正常
        - 根据[test-resources.yaml](test-resources.yaml)测试
            ```bash
            kubectl apply -f test-resources.yaml
            # 查看状态
            kubectl describe certificate -n cert-manager-test
            # 可以看到证书
            kubectl describe -n cert-manager-test secret selfsigned-cert-tls
            # 查看证书过期时间
            kubectl get -n cert-manager-test secrets selfsigned-cert-tls -o jsonpath="{.data.ca\.crt}" | base64 -d | openssl x509 -enddate -noout
            ```
            ![](picture/test-cert-manager.png)
        - 清除
            ```bash
            kubectl delete -f test-resources.yaml
            ```
## 使用

安装完`cert-manager`后必须创建`Issuer`或者`ClusterIssuer`资源作为证书认证. 
请参考[官方指导](https://cert-manager.io/docs/configuration/) 配置不同证书来源的`Issuer`

### 创建issuer和clusterissuer
**`注意: 未实践`** 生产环境中推荐使用`Let‘s Encrypt`等证书签发机构签发证书. 以[acme](https://cert-manager.io/docs/configuration/acme/) 为例:
- 当你创建`ACME Issuer`时, cert-manager会生成一个私钥, 其用acme server验证你的信息.
    - 使用[cert-manager-secret.yaml](cert-manager-secret.yaml)
    ```bash
    kubectl apply -f cert-manager-secret.yaml
    ```
- 

#### SelfSigned方式(实际使用)
使用[official.yaml](official.yaml), [参考](https://cert-manager.io/docs/configuration/selfsigned/)
```bash
kubectl apply -f official.yaml
```

#### CA 证书方式(未采用)
我们使用[ca](https://cert-manager.io/docs/configuration/ca/) 创建`Issuer`等资源
- 生产环境中推荐使用[letsencrypt](https://letsencrypt.org/) 等机构签发的免费ssl证书  **`未能验证是否正确`**
    - 快速(以acme为例): 使用[cert-manager-secret.yaml](cert-manager-secret.yaml) , [参考](https://blog.csdn.net/weixin_44692256/article/details/108274385)
        ```bash
        kubectl apply -f cert-manager-secret.yaml
        # 查看clusterissuer
        kubectl get clusterissuers traefik-issuer -n cert-manager-use -o wide
        ```
- 测试部署当中, 这里使用k8s集群的证书使用
    ![](picture/kubernetes-ca.png)
        ```bash
        cp /etc/kubernetes/pki/ca.crt ./
        cp /etc/kubernetes/pki/ca.key ./
        ```
        ![](picture/ca-use.png)
    - 也可手动创建证书, [参考](https://kubernetes.io/zh/docs/tasks/administer-cluster/certificates/#openssl)  (不采用)
- 创建secret: `ca-key-pair`及命名空间: `cert-manager-use`
    - 方式1: 使用[official-ca.yaml](official-ca.yaml) (参考[test-resources.yaml](test-resources.yaml) 即可创建自定义证书)
        - 将`ca.crt`、`ca.key`内的内容分别拷入至`data`. **注意要使用`64`位读取**, 
            ```bash
            cat ca.crt | base64 -w0
            cat ca.key | base64 -w0
            ```
            ![](picture/copy-ca-to-yaml.png)
        ```bash
        kubectl apply -f official.yaml
        ```
- 创建完成后查看:
    ```bash
    kubectl describe secret ca-key-pair -n cert-manager-use
    # 查看证书过期时间
    openssl x509 -in tls.crt -noout -dates
    ```
  
### 使用issuer等生成Certificate证书
[参考](https://cert-manager.io/docs/usage/)


## 其他知识
详细了解(不采用): 在颁发证书前, 你必须在集群中配置至少一个`Issuer` 或者`ClusterIssuer`(可跨命名空间使用) resource
- 创建`Issuer`, [参考](https://cert-manager.io/docs/concepts/issuer/)

[设置向导](https://cert-manager.io/docs/configuration/) 会教你如何配置颁发证书
- [cert-manager使用ca](https://cert-manager.io/docs/configuration/ca/)

### kubectl 插件
[cert-manager插件](https://cert-manager.io/docs/usage/kubectl-plugin/): 可以使用该插件在
kubernetes集群中管理 cert-manager资源

## 参考
- [cert-manager+Let‘s Encrypt自动证书签发](https://blog.csdn.net/ai524719755/article/details/116712931)
- [k8s中使用cert-manager玩转证书-CA方式](https://cloud.tencent.com/developer/article/1402451)