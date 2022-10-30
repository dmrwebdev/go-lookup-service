package server

import (
	"golookupservice/pkg/lookup"
	"net/http"
	"strings"
)

const lookupRoute = "lookup"

// Declared routes for the server. Add additional routes here
func handleRoutes(mux *http.ServeMux, s *httpServer) {
	mux.HandleFunc("/", s.getRoot)
	mux.HandleFunc(lookupRoute, s.getLookup)
}

// Handler for root route
func (s *httpServer) getRoot(w http.ResponseWriter, r *http.Request) {
	l := lookup.NewLookup(w, s.cache)
	loc := map[string]string{
		"city": r.URL.Query().Get("city"),
		"country": r.URL.Query().Get("country"),
	}

	if loc["city"] == "" && loc["country"] == "" {
		l.GetAll()
	} else {
		l.GetFiltered(loc)
	}
} 

// Handler for lookup route
func (s *httpServer) getLookup(w http.ResponseWriter, r *http.Request) {
	ip := strings.TrimPrefix(r.URL.Path, lookupRoute)
	l := lookup.NewLookup(w, s.cache)
	l.GetIp(ip)
}