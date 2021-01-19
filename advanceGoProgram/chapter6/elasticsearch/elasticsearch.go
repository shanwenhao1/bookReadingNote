package elasticsearch

// 选用 elastic 版本时
// 注意与自己使用的 elasticsearch 要对应
import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v3"
)

var esClient *elastic.Client

func InitElasticsearchClient(host string, port string) {
	var err error
	esClient, err = elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("http://%s:%s", host, port)),
		elastic.SetMaxRetries(3),
	)

	if err != nil {
		// log error
	}
}

// 插入
func insertDocument(db string, table string, obj map[string]interface{}) {

	//id := obj["id"]

	var indexName, typeName string
	// 数据库中的 database/table 概念，可以简单映射到 es 的 index 和 type
	// 不过需要注意，因为 es 中的 _type 本质上只是 document 的一个字段
	// 所以单个 index 内容过多会导致性能问题
	// 在新版本中 type 已经废弃
	// 为了让不同表的数据落入不同的 index，这里我们用 table+name 作为 index 的名字
	indexName = fmt.Sprintf("%v_%v", db, table)
	typeName = table

	// 正常情况
	res, err := esClient.Index().Index(indexName).Type(typeName).Id("id").BodyJson(obj).Do()
	if err != nil {
		// handle error
	} else {
		// insert success
	}
	fmt.Println("TODO handle res", res)
}

// 获取
func query(indexName string, typeName string) (*elastic.SearchResult, error) {
	// 通过 bool must 和 bool should 添加 bool 查询条件
	q := elastic.NewBoolQuery().Must(elastic.NewMatchPhraseQuery("id", 1),
		elastic.NewBoolQuery().Must(elastic.NewMatchPhraseQuery("male", "m")))

	q = q.Should(
		elastic.NewMatchPhraseQuery("name", "alex"),
		elastic.NewMatchPhraseQuery("name", "xargin"),
	)

	searchService := esClient.Search(indexName).Type(typeName)
	res, err := searchService.Query(q).Do()
	if err != nil {
		// log error
		return nil, err
	}

	return res, nil
}

// 删除
func deleteDocument(indexName string, typeName string, obj map[string]interface{}) {
	//id := obj["id"]

	res, err := esClient.Delete().Index(indexName).Type(typeName).Id("id").Do()
	if err != nil {
		// handle error
	} else {
		// delete success
	}
	fmt.Println("TODO handle res", res)
}
