package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/scliang-strive/webServerTools/config"
	"github.com/sirupsen/logrus"
)

var esClient *elasticsearch.Client

type EsClient struct {
	Client *elasticsearch.Client
}

func InitEsClient(cfg *config.Elasticsearch) {
	c := elasticsearch.Config{
		Addresses: cfg.Host,
		Username:  cfg.Username,
		Password:  cfg.Password,
		CloudID:   cfg.CloudId,
		APIKey:    cfg.APIKey,
	}
	client, err := elasticsearch.NewClient(c)
	if err != nil {
		logrus.Errorf("init elasticsearch failed. [ERROR]: %s", err.Error())
		return
	}
	esClient = client
}

func GetEs() *EsClient {
	return &EsClient{
		Client: esClient,
	}
}

// Index 在索引中创建或更新文档
// 索引不存在的情况下，会自动创建索引
func (es *EsClient) Index(index, docType string, doc map[string]interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(doc)
	if err != nil {
		logrus.Errorf("[ERROR]: error encoding doc, %s", err.Error())
		return err
	}

	res, err := es.Client.Index(index, &buf, es.Client.Index.WithDocumentType(docType))
	if err != nil {
		logrus.Errorf("[ERROR]: Index response failed, %s", err.Error())
	}
	defer func() {
		_ = res.Body.Close()
	}()
	return nil
}

// Create 添加文档需要指定_id，_id已存在返回409
func (es *EsClient) Create(index, id, docType string, doc map[string]interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(doc)
	if err != nil {
		logrus.Errorf("[ERROR]: encoding doc failed, %s", err.Error())
		return err
	}
	res, err := es.Client.Create(index, id, &buf, es.Client.Create.WithDocumentType(docType))
	if err != nil {
		logrus.Errorf("[ERROR]: create response failed, %s", err.Error())
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	return nil
}

// DeleteByQuery 通过匹配条件删除文档
func (es *EsClient) DeleteByQuery(index []string, query map[string]interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(query)
	if err != nil {
		logrus.Errorf("[ERROR]: encoding query failed, %s", err.Error())
		return err
	}
	res, err := es.Client.DeleteByQuery(index, &buf)
	if err != nil {
		logrus.Errorf("[ERROR]: delete by query response failed, %s", err.Error())
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	return nil
}

// Delete 通过_id删除文档
func (es *EsClient) Delete(index, id string) error {
	res, err := es.Client.Delete(index, id)
	if err != nil {
		logrus.Errorf("[ERROR]: delete by id response failed, %s", err.Error())
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	return nil
}

// Search 搜索
func (es *EsClient) Search(index string, query map[string]interface{}) (map[string]interface{}, error) {
	_, err := es.Client.Info()
	if err != nil {
		logrus.Errorf("[ERROR]: getting response failed, %s", err.Error())
		return nil, err
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(query)
	if err != nil {
		logrus.Errorf("[ERROR]: encoding query failed, %s", err.Error())
		return nil, err
	}
	res, err := es.Client.Search(
		es.Client.Search.WithContext(context.Background()),
		es.Client.Search.WithIndex(index),
		es.Client.Search.WithBody(&buf),
		es.Client.Search.WithTrackTotalHits(true),
		es.Client.Search.WithPretty(),
	)
	if err != nil {
		logrus.Errorf("[ERROR]: getting response failed, %s", err.Error())
		return nil, err
	}
	var e map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		logrus.Errorf("[ERROR]: decode reponse failed, %s",err.Error())
		return nil,err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	return e, nil
}

// Get 通过id获取文档
func (es *EsClient) Get(index, id string) (map[string]interface{}, error) {
	res, err := es.Client.Get(index, id)
	if err != nil {
		logrus.Errorf("[ERROR]: get response failed, %s", err.Error())
		return nil, err
	}
	var e map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		logrus.Errorf("[ERROR]: decode reponse failed, %s",err.Error())
		return nil,err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	return e, nil
}

// Update 通过_id更新文档
func (es *EsClient) Update(index, id, docType string, doc map[string]interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(doc)
	if err != nil {
		logrus.Errorf("[ERROR]: encoding query failed, %s", err.Error())
		return err
	}
	res, err := es.Client.Update(index, id, &buf, es.Client.Update.WithDocumentType(docType))
	if err != nil {
		logrus.Errorf("[ERROR]: update documents failed, %s", err.Error())
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	return nil
}

// UpdateByQuery 通过匹配条件更新文档
func (es *EsClient) UpdateByQuery(index []string, docType string, doc map[string]interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(doc)
	if err != nil {
		logrus.Errorf("[ERROR]: encoding query failed, %s", err.Error())
		return err
	}
	res, err := es.Client.UpdateByQuery(
		index,
		es.Client.UpdateByQuery.WithDocumentType(docType),
		es.Client.UpdateByQuery.WithContext(context.Background()),
		es.Client.UpdateByQuery.WithBody(&buf),
		es.Client.UpdateByQuery.WithPretty(),
	)
	if err != nil {
		logrus.Errorf("[ERROR]: udpate by quesy failed, %s", err.Error())
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	return nil
}
