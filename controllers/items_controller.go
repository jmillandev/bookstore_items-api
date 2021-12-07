package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jmillandev/bookstore_items-api/domains/items"
	"github.com/jmillandev/bookstore_items-api/services"
	"github.com/jmillandev/bookstore_items-api/utils/http_utils"
	"github.com/jmillandev/bookstore_oauth-go/oauth"
	"github.com/jmillandev/bookstore_utils-go/rest_errors"
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
		http_utils.ResponseJsonError(w, err)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		err := rest_errors.NewUnauthorizedError()
		http_utils.ResponseJsonError(w, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.ResponseJsonError(w, respErr)
		return
	}
	defer r.Body.Close()

	var item items.Item
	if err := json.Unmarshal(requestBody, &item); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.ResponseJsonError(w, respErr)
		return
	}

	item.Seller = sellerId

	result, createErr := services.ItemsService.Create(item)
	if createErr != nil {
		http_utils.ResponseJsonError(w, createErr)
		return
	}
	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {}
