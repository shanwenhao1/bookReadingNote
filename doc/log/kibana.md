# Kibana安装及使用

## kibana配置
[官方参考](https://www.elastic.co/guide/en/kibana/7.10/settings.html)

## docker方式安装
```bash
# docker run --link YOUR_ELASTICSEARCH_CONTAINER_NAME_OR_ID:elasticsearch -p 5601:5601 docker.elastic.co/kibana/kibana:7.10.1
# kibana版本号最好与elasticsearch版本号一致
docker run -d --link es01:elasticsearch --net elasticsearch_elastic -p 5601:5601 --name kibana kibana:7.10.1
# 可用docker inspec es01查看网络

-----或者------
# 连接至es集群中的单个节点
docker run -d -e ELASTICSEARCH_URL=http://192.168.1.89:9200 -p 5601:5601 --name kibana kibana:7.10.1
```

## docker-compose方式安装

使用[kibana.yml](yml/kibana/kb-config/kibana.yml)配置文件(该配置文件要置入执行命令的目录下), 
用[docker-compose.yml](yml/kibana/docker-compose.yml)部署
```bash
docker-compose up -d
```

## 参考
- [官方指导](https://www.elastic.co/guide/en/kibana/7.10/docker.html)