package server

import (
	"errors"
	"fmt"
	"golookupservice/pkg/lookup"
	"net/http"
	"time"
)

type httpServer struct {
	cache *lookup.LocalCache
}

// Create new HTTP server
func NewServer(port int) *http.Server{
	s := &httpServer{
		cache: lookup.NewLocalCache(30*time.Minute),
	}

	mux := http.NewServeMux()
	handleRoutes(mux, s)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: mux,
	}
}

// Start a HTTP server
func StartServer(s *http.Server) {
		err := s.ListenAndServe()
		
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
}