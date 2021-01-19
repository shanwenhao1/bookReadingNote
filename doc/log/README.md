# EFk 日志收集系统搭建及使用

EFK指的是[Elasticsearch](https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html) + 
[Fluentd]() + [Kibana]()
- elasticsearch: 是一个天然的分布式文档存储中间件, 它使用倒排索引的数据结构, 支持快速全文搜索.
- kibana: 集成了`elasticsearch`中的安全、监控和管理功能, 用作web端管理界面.

## 搭建
- [elasticsearch集群搭建](elasticsearch.md)