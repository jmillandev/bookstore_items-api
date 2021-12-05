package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmillandev/bookstore_items-api/clients/elasticsearch"
	"github.com/jmillandev/bookstore_utils-go/logger"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	MapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Server starting on port 8888...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
