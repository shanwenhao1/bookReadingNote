package main

import (
	"bookReadingNote/infra/redis"
	"bookReadingNote/infra/tool/file/xmlFile"
	"fmt"
)

/*
	服务启动前的准备, 包括redis连接
*/
func InitServer() {
	/*
		---------------------------本地配置文件方式加载配置----------------------------------
	*/
	// 加载redis 配置文件
	xmlF, err := xmlFile.XmlParse("config/redisConfig.xml", new(xmlFile.RDConfig))
	if err != nil {
		panic(err)
	}
	fmt.Println(xmlF.(*xmlFile.RDConfig))
	// redis 连接初始化
	redis.InitRedis(xmlF)
}

func main() {
	InitServer()
	redis.ExampleClient()
}
