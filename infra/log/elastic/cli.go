package elasticbase

import (
	"bookReadingNote/infra/log/elastic/esconf"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

/*
ElsConfig elasticsearch 连接config
*/
type ElsConfig struct {
	Address  []string `json:"address"`  // connect address, example: "http://192.168.0.170:9200"
	Username string   `json:"username"` // elastic username, example: elastic
	Password string   `json:"password"` // elastic password, example: test123
}

/*
ElsData elastic操作
*/
type ElsData struct {
	cli *elasticsearch.Client
}

// newData generate a new elasticsearch handle client
func newData(cli *elasticsearch.Client) *ElsData {
	var newD = new(ElsData)
	newD.cli = cli
	return newD
}

// NewElsData generate a new elasticsearch handle client
func NewElsData(esCfg *esconf.EsConfig) *ElsData {
	if esCfg == nil {
		return nil
	}
	cfg := elasticsearch.Config{
		Addresses: esCfg.Address,
		Username:  *esCfg.Username,
		Password:  *esCfg.Password,
	}
	// 如果连接失败则返回nil, 不panic
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("create elasticsearch client connect failed, err: %v", err)
		return nil
	}
	return newData(es)
}
