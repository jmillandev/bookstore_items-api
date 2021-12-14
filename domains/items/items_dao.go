package items

import (
	"errors"

	"github.com/jmillandev/bookstore_items-api/clients/elasticsearch"
	"github.com/jmillandev/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when try to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() *rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError("no item found with given id")
		}
		return rest_errors.NewInternalServerError("error when try to get item", errors.New("database error"))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when try to parse item", err)
	}
	i.Id = itemId
	return nil
}
