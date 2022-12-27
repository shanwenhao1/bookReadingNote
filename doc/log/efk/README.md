# efk 日志分析系统

[elastic官方 docker](https://www.docker.elastic.co/)

EFK指的是[Elasticsearch](https://www.elastic.co/guide/en/elasticsearch/reference/index.html) +
[Filebeat](https://www.elastic.co/guide/en/beats/filebeat/index.html) +
[Kibana](https://www.elastic.co/guide/en/kibana/index.html)
- elasticsearch: 是一个天然的分布式文档存储中间件, 它使用倒排索引的数据结构, 支持快速全文搜索.
- kibana: 集成了`elasticsearch`中的安全、监控和管理功能, 用作web端管理界面.

## 服务搭建

使用 [docker-compose.yml](docker-compose.yml) 安装服务([docker-compose命令安装](../centos/tool.md#docker-compose))
- 启动前的准备
    ```bash
    # 创建elasticsearch、filebeat 数据挂载目录并赋予权限(否则会导致elastic运行失败)
    mkdir esdata01 filebeat
    chmod 777 esdata01 filebeat
    ```
    - 将`filebeat`的配置文件[filebeat.yml](filebeat/filebeat-example.yml)置入filebeat目录下
- 启动服务
```bash
# 一键启动服务
sudo docker compose up -d
```
- `curl -X GET "http://localhost:9200" -u elastic:abc123`

### 访问
访问`http://192.168.0.170:5601`
- 配置token
  - `docker exec -it es01 bin/bash`
    - `bin/elasticsearch-create-enrollment-token --scope kibana`
- 配置验证
  - `docker exec -it kibana /bin/bash`
    - `bin/kibana-verification-code`

## 数据清理
[文档](esDataClean/README.md)



## 参考
- [官方docker-compose](https://github.com/elastic/elasticsearch/blob/master/docs/reference/setup/install/docker/docker-compose.yml)
- [其他博客efk参考](https://www.elastic.co/guide/en/kibana/7.10/settings.html)
- [https://github.com/MihowBogucki/local-docker-efk](https://github.com/MihowBogucki/local-docker-efk)

## 其他
许可证安装
```bash
# 上传许可证(暂时没用)
curl -XPUT -u elastic 'http://127.0.0.1:9200/_xpack/license?acknowledge=true' -H "Content-Type: application/json" -d @ license.json

# 开启basic license 使用
curl -XPOST -u elastic 'http://127.0.0.1:9200/_license/start_basic?acknowledge=true'

# 设置最大切片数, 可用命令curl  -u elastic "localhost:9200/_cluster/settings?pretty" 查看, https://www.elastic.co/guide/en/elasticsearch/reference/7.17/transient-settings-migration-guide.html
PUT /_cluster/settings
{
  "persistent": {
    "cluster": {
      "max_shards_per_node":10000
    }
  }
}
```