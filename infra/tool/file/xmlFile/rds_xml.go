package xmlFile

import "time"

var RDSCfg RDConfig

/*
	redis xml配置文件对应 struct
*/
type RDConfig struct {
	XmlHandle
	RdAddr      string        `xml:"redis_addr"`
	Pass        string        `xml:"redis_password"`
	DbNum       int           `xml:"redis_db_num"` // 采用redis哪个数据库
	PoolSize    int           `xml:"redis_pool_size"`
	IdleTimeout time.Duration `xml:"redis_idle_timeout"` // 空闲连接失效, 使用时需 * time.second等
}
