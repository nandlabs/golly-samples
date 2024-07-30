package server

import (
	"log"
	"net/http"
	"time"

	"github.com/nandlabs/golly-samples/turbo/server-routing/handlers"
	"github.com/nandlabs/golly-samples/turbo/server-routing/store"
	"oss.nandlabs.io/golly/turbo"
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
	router := turbo.NewRouter()

	// register routes
	router.Get("/api/v1/users", handlers.GetUsers(s.store))
	router.Post("/api/v1/users", handlers.AddUser(s.store))
	router.Put("/api/v1/users/:id", handlers.UpdateUser(router, s.store))
	router.Delete("/api/v1/users/:id", handlers.DeleteUser(router, s.store))

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	// to start the server, invoke the ListenAndServe method
	if err := srv.ListenAndServe(); err != nil {
		// We will use our inbuilt logger in future examples
		log.Fatalln(err)
	}
}
