package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

const KeyServerAddr = "serverAddr"

func StartServer(s http.Server) {
		err := s.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
}

func StartMultipleServers(s http.Server, c context.CancelFunc) {
		err := s.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		c()
}

func GenerateServer(port int, mux *http.ServeMux, ctx context.Context) *http.Server{
	return &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, KeyServerAddr, l.Addr().String())
			return ctx
		},
	}
}