package xmlFile

/*
	mysql xml 配置文件对应struct
*/
type MysqlConfig struct {
	XmlHandle
	DbName     string `xml:"db_name"`
	DbUser     string `xml:"db_user"`
	DbPwd      string `xml:"db_pwd"`
	DbUrl      string `xml:"db_url"`
	DbMaxConn  int    `xml:"db_max_conn"`
	DbMaxIdle  int    `xml:"db_max_idle"`
	DbLogModel bool   `xml:"db_log_model"`
}
