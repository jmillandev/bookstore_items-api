package app

import (
	"net/http"

	"github.com/jmillandev/bookstore_items-api/controllers"
)

func MapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
