package main

import (
	"log"
	"net/http"
	"time"

	"oss.nandlabs.io/golly/turbo"
)

func main() {
	router := turbo.NewRouter()
	router.Get("/api/v1/healthz", healthCheck)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server is up and running"))
	w.WriteHeader(http.StatusOK)
}
