#!/bin/bash

## 清理30天之前的数据, 示例:2022.07.28
#his_date=$(date -d "-1 day" +"%Y.%m.%d")
#
#echo $his_date
#
# curl -XDELETE "http://elastic:abc123@127.0.0.1:9200/*-$(his_date)?pretty"

# 其中elastic:abc123为账号密码, 清理30天之前的数据
curl -XPOST "http://elastic:abc123@127.0.0.1:9200/*/_delete_by_query" -H 'Content-Type: application/json' -d'{
  "query": {
    "range": {
      "@timestamp": {
        "lt": "now-30d",
        "format": "epoch_millis"
      }
    }
  }
}'