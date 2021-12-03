package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	MapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
