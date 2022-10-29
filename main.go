package main

import (
	"context"
	"golookupservice/pkg/server"
	"net/http"
)

func main() {
	// Use custom multiplexer instead of default
	mux := http.NewServeMux()
	// Assign functions to routes
	mux.HandleFunc("/", server.GetRoot)
	mux.HandleFunc("/lookup/", server.GetLookup)

	// For single server instance
	ctx := context.Background()
	s1 := server.GenerateServer(3333, mux, ctx)
	server.StartServer(*s1)

// If additional servers are needed...	
//ctx, cancelCtx := context.WithCancel(context.Background())
//s1 := server.GenerateServer(3333, mux, ctx)
//s2 := server.GenerateServer(44499, mux, ctx)
//go server.StartMultipleServers(*s1, cancelCtx)
//go server.StartMultipleServers(*s2, cancelCtx)

//<-ctx.Done()

}