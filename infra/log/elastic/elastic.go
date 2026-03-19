package elasticbase

import (
	"bookReadingNote/infra/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
)

/*
ElasticTool client tool for elasticsearch
*/
type ElasticTool interface {
	Create(index string, data map[string]interface{}) (map[string]interface{}, error)
	Query(index string, id string) (map[string]interface{}, error)
}

/*
Create 添加新的文档

	index: elastic document index
	data: 添加的data数据
*/
func (cli ElsData) Create(index string, data map[string]interface{}) (map[string]interface{}, error) {
	var (
		err  error
		buf  bytes.Buffer
		resM map[string]interface{}
	)
	// 默认document timestamp
	data = cli.defaultTime(data)
	err = json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return nil, err
	}
	// 创建文档
	res, err := cli.cli.Index(index, &buf, cli.cli.Index.WithDocumentType("_doc"))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 201 {
		return nil, errors.New(fmt.Sprintf("create document failed, resp: %v", res))
	}
	// 解析返回数据
	resM, err = cli.parseResponse(res)
	if err != nil {
		return nil, err
	}
	return resM, nil
}

/*
Query 查询最新的文档

	index: elastic document index
	id: elastic document id
*/
func (cli ElsData) Query(index string, id string) (map[string]interface{}, error) {
	var (
		err  error
		resM map[string]interface{}
	)
	// 查询文档
	res, err := cli.cli.Get(index, id)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("query document failed, resp: %v", res))
	}
	// 解析返回数据
	resM, err = cli.parseResponse(res)
	if err != nil {
		return nil, err
	}
	return resM, nil
}

// 默认时间数据
func (cli ElsData) defaultTime(data map[string]interface{}) map[string]interface{} {
	data["@timestamp"] = utils.GetCurTimeUtc()
	return data
}

func (cli ElsData) parseResponse(res *esapi.Response) (map[string]interface{}, error) {
	var resD map[string]interface{}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resD)
	if err != nil {
		return nil, err
	}
	return resD, nil
}
