package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/jmillandev/bookstore_utils-go/logger"
	"github.com/olivere/elastic"
)

var Client esClientInterface = &esClient{}

type esClientInterface interface {
	Index(string, interface{}) (*elastic.IndexResponse, error)
	setClient(*elastic.Client)
}

type esClient struct {
	client *elastic.Client
}

func Init() {

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) Index(index string, data interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index("Items").
		BodyJson(data).
		Do(ctx)
	if err != nil {
		logger.Error(
			fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}
