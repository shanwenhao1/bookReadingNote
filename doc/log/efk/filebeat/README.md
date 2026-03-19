# filebeat

`filebeat`配置请参考官方 [filebeat-reference-yml](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-reference-yml.html)


```bash
docker run -d --net=host --restart=always --name=filebeat --user=root -v /root/efk/filebeat/filebeat.yml:/usr/share/filebeat/filebeat.yml:ro \
  -v /root/ibserver/user/logs:/root/logs \
  docker.elastic.co/beats/filebeat:7.17.1 filebeat -e
docker stop filebeat && docker rm filebeat


docker run --net=host --restart=always --name=filebeat --user=root -v /root/efk/filebeat/filebeat.yml:/usr/share/filebeat/filebeat.yml:ro \
  -v /root/ibserver/user/logs:/root/logs \
  docker.elastic.co/beats/filebeat:7.17.1 filebeat -e
```

## filebeat配置
[filebeat可优化配置](https://blog.csdn.net/qq_41926119/article/details/104549808)