# Redis部署

## 使用Docker方式部署

[参考](https://www.cnblogs.com/934827624-qq-com/p/10175478.html)

[redis.conf](redis.conf)
```bash
# 注释掉配置文件中的 daemonize yes
# daemonize yes
```

运行redis镜像
```bash
# 下载镜像
#docker pull redis:3.2
docker pull redis:latest
# 创建redis持久化目录 ~/redis/data
mkdir -p /root/redis /root/redis/data /root/redis/conf
# 将redis.conf导入至 '/root/redis/conf' 目录下, 确保redis.conf的权限为 '-rw-r--r--' 

# 启动redis, 持久化到磁盘/root/redis/data
# docker run -d  -p 6379:6379 --name my_redis -v /root/redis/data:/data redis:latest  --appendonly yes --requirepass "PASSWD"

# 启动redis, 持久化到磁盘/root/redis/data(以配置文件的方式启动), 密码是abc123
docker run -d --privileged=true --restart=always -p 6379:6379 --name my_redis -v /root/redis/data:/data -v /root/redis/conf/redis.conf:/etc/redis/redis.conf -d redis redis-server /etc/redis/redis.conf --requirepass abc123
docker run -d --privileged=true --net=host --restart=always -p 6379:6379 --name my_redis --requirepass abc123 redis
```

### 访问

- 本机访问容器redis服务可使用
```bash
# localhost:6379或127.0.0.1:6379
# 宿主机内网同内网地址 192.168.1.89:6379
```
- 容器依赖另外一个容器的服务
```bash
# 必须使用宿主机的内网或者公网ip访问
```

## K8s部署方式
TODO 待补充

## Redis数据迁移

- 修改`redis.conf`文件, 并重启服务`docker restart my_redis`
```bash
# 修改`client-output-buffer-limit`
# client-output-buffer-limit slave 256mb 64mb 60
client-output-buffer-limit slave 0 0 0 
```

- 参考[阿里云文档](https://help.aliyun.com/document_detail/117311.html?spm=5176.10695662.1996646101.searchclickresult.4ce36541rNxoq0)
进行数据迁移
    - 下载[redis-shake](https://github.com/alibaba/RedisShake/releases?spm=a2c4g.11186623.2.12.4dcd6f10KeICfP)
    - 将`redis-shake.tar.gz`置入目录`/root/redis/redis-shake`并解压
        ```bash
        tar -xvf redis-shake-1.6.16.tar.gz
        ```
    - 更改redis-shake配置文件`redis-shake.conf`, 修改后[redis-shake.conf](redis-shake.conf)(从阿里云redis迁移数据至
    自建redis服务器)
    - 运行命令进行迁移
        ```bash
        ./redis-shake -type=rump -conf=redis-shake.conf
        ```