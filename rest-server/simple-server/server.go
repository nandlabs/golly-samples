package main

import (
	"net/http"

	"oss.nandlabs.io/golly/lifecycle"
	"oss.nandlabs.io/golly/rest/server"
)

func main() {
	// register the router by creating the server object
	restServer, err := server.Default()
	if err != nil {
		panic(err)
	}

	// Add path prefix if you want
	restServer.Opts().PathPrefix = "/api/v1"

	// add a GET endpoint to your server by providing the "endpoint" and handler function
	restServer.Get("/healthz", healthCheck)

	// create the http.Server object and register the router as Handler
	// provide the necessary configurations such as PORT, ReadTimeout, WriteTimeout...
	manager := lifecycle.NewSimpleComponentManager()

	// Register the server
	manager.Register(restServer)

	// start the server
	manager.StartAndWait()
}

func healthCheck(ctx server.Context) {
	ctx.HttpResWriter().Write([]byte("server is up and running"))
	ctx.HttpResWriter().WriteHeader(http.StatusOK)
}
