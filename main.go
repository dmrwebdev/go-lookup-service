package main

import (
	"context"
	"errors"
	"fmt"
	"gohttpserver/pkg/lookup"
	"gohttpserver/pkg/server"
	"io"
	"net"
	"net/http"
	"strings"
)



func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	city := r.URL.Query().Get("city")
	country := r.URL.Query().Get("country")
	
	fmt.Printf("%s: got /request. city=%s, country=%s\n",
	ctx.Value(server.KeyServerAddr),	city,	country)
	io.WriteString(w, "Lookup!")
} 

func getLookup(w http.ResponseWriter, r *http.Request) {
	ip := strings.TrimPrefix(r.URL.Path, "/lookup/")
	lookup.GetIp(ip, w)
}



func main() {
	// Use custom multiplexer instead of default
	mux := http.NewServeMux()
	// Assign functions to routes
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/lookup/", getLookup)

	ctx := context.Background()
	serverOne := &http.Server{
		Addr: ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, server.KeyServerAddr, l.Addr().String())
			return ctx
		},
	}

	err := serverOne.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server one closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server one: %s\n", err)
	}

// If additional servers are needed...	
//ctx, cancelCtx := context.WithCancel(context.Background())
//serverOne := server.GenerateServer(3333, mux, ctx,)
//serverTwo := server.GenerateServer(44499, mux, ctx,)
//go server.StartServer(*serverOne, cancelCtx)
//go server.StartServer(*serverTwo, cancelCtx)

//<-ctx.Done()

}