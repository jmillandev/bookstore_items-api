package controllers

import (
	"fmt"
	"net/http"

	"github.com/jmillandev/bookstore_items-api/domains/items"
	"github.com/jmillandev/bookstore_items-api/services"
	"github.com/jmillandev/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: Return error to the caller
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO: Reutrn error json to the user.
		return
	}

	fmt.Println(result)
	// TODO: Return create item as json with HTTP status 201 - Created
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {}
