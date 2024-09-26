package server

import (
	"github.com/nandlabs/golly-samples/rest-server/server-routing/handlers"
	"github.com/nandlabs/golly-samples/rest-server/server-routing/store"
	"oss.nandlabs.io/golly/lifecycle"
	"oss.nandlabs.io/golly/rest/server"
)

type Server struct {
	store *store.Store
}

func NewServer() *Server {
	return &Server{
		store: store.NewStore(),
	}
}

func (s *Server) Start() {

	// register the router by creating the server object
	restServer, err := server.Default()
	if err != nil {
		panic(err)
	}

	// Add path prefix if you want
	restServer.Opts().PathPrefix = "/api/v1"

	// register routes
	restServer.Get("/users", handlers.GetUsers)
	restServer.Post("/users", handlers.AddUser)
	restServer.Put("/users/:id", handlers.UpdateUser)
	restServer.Delete("/users/:id", handlers.DeleteUser)

	// create the http.Server object and register the router as Handler
	// provide the necessary configurations such as PORT, ReadTimeout, WriteTimeout...
	manager := lifecycle.NewSimpleComponentManager()

	// Register the server
	manager.Register(restServer)

	// start the server
	manager.StartAndWait()
}
