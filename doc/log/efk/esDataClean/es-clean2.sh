#!/bin/bash

# 清理app ser日志
curl -XPOST "http://elastic:abc123@127.0.0.1:9200/appser*/_delete_by_query" -H 'Content-Type: application/json' -d'{
  "query": {
    "range": {
      "@timestamp": {
        "lt": "now-30d",
        "format": "epoch_millis"
      }
    }
  }
}'
# 清理admin ser日志
curl -XPOST "http://elastic:abc123@127.0.0.1:9200/adminser*/_delete_by_query" -H 'Content-Type: application/json' -d'{
  "query": {
    "range": {
      "@timestamp": {
        "lt": "now-30d",
        "format": "epoch_millis"
      }
    }
  }
}'
# TODO 添加剩余的清除命令