package main

import (
	"log"
	"net/http"
	"time"

	"oss.nandlabs.io/golly/turbo"
)

func main() {
	// register the router by creating the turbo object
	router := turbo.NewRouter()

	// add a GET endpoint to your server by providing the "endpoint" and handler function
	router.Get("/api/v1/healthz", healthCheck)

	// create the http.Server object and register the router as Handler
	// provide the necessary configurations such as PORT, ReadTimeout, WriteTimeout...
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	// to start the server, invoke the ListenAndServe method
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server is up and running"))
	w.WriteHeader(http.StatusOK)
}
