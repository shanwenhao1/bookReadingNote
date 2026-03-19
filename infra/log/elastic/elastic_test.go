package elasticbase

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testIndex = "test"
)

func testInit() *ElsData {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.0.170:9200",
		},
		Username: "elastic",
		Password: "abc123",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	elsD := newData(es)
	return elsD
}

func TestElsData_Create(t *testing.T) {
	a := assert.New(t)
	cli := testInit()

	testStr := "test123"
	data := map[string]interface{}{
		"msg": testStr,
	}
	resMap, err := cli.Create(testIndex, data)
	a.Equal(nil, err, "create document failed, err: %v", err)
	a.NotEqual("", resMap["_id"], "create document failed, wrong id")
	//fmt.Println("----", resMap)

	// 查询并校验生成是否成功生成
	qMap, err := cli.Query(testIndex, resMap["_id"].(string))
	//fmt.Println("----query", qMap)
	a.Equal(nil, err, "query document failed, err: %v", err)
	_source := qMap["_source"].(map[string]interface{})
	a.Equal(testStr, _source["msg"].(string), "query document failed, wrong msg")
}
