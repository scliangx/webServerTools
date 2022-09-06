package elasticsearch

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"testing"
)

func TestES(t *testing.T) {
	defaultClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		fmt.Println("default error", err.Error())
	}
	fmt.Println("-----------------------------")
	client := EsClient{
		defaultClient,
	}
	fmt.Println("client: ", client)
	doc := map[string]interface{}{"web1": "web10doc"}
	err = client.Index("demo", "doc", doc)
	if err != nil {
		fmt.Println("error: ",err.Error())
	}
}
