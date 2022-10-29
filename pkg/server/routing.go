package server

import (
	"fmt"
	"golookupservice/pkg/lookup"
	"io"
	"net/http"
	"strings"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	city := r.URL.Query().Get("city")
	country := r.URL.Query().Get("country")
	
	fmt.Printf("%s: got /request. city=%s, country=%s\n",
	ctx.Value(KeyServerAddr),	city,	country)
	io.WriteString(w, "Lookup!")
} 

func GetLookup(w http.ResponseWriter, r *http.Request) {
	ip := strings.TrimPrefix(r.URL.Path, "/lookup/")
	lookup.GetIp(ip, w)
}