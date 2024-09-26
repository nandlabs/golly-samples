package main

import "github.com/nandlabs/golly-samples/rest-server/server-routing/server"

func main() {
	srv := server.NewServer()
	srv.Start()
}
