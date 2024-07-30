package main

import "github.com/nandlabs/golly-samples/turbo/server-routing/server"

func main() {
	srv := server.NewServer()
	srv.Start()
}
