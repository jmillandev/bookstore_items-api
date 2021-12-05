package elasticsearch

import (
	"context"
	"time"

	"github.com/olivere/elastic"
)

var Client esClientInterface = &esClient{}

type esClientInterface interface {
	Index(interface{}) (*elastic.IndexResponse, error)
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

func (c *esClient) Index(data interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().Do(ctx)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}
