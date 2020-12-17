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
	var rdCfg = new(xmlFile.RDConfig)
	err := xmlFile.XmlParse("config/redisConfig.xml", new(xmlFile.RDConfig), rdCfg)
	if err != nil {
		panic(err)
	}
	//fmt.Println(rdCfg.(*xmlFile.RDConfig))
	// redis 连接初始化
	redis.InitRedis(rdCfg)
}

func redisTest() {
	rd := new(redis.RDSHandle)
	err := rd.SetNX("test_2", 2)
	fmt.Println(err)
}

func main() {
	InitServer()
	redisTest()
}
