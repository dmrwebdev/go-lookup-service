package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

func ipLookup() {

}


func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	city := r.URL.Query().Get("city")
	country := r.URL.Query().Get("country")
	
	fmt.Printf("%s: got /request. city=%s, country=%s\n",
	ctx.Value(keyServerAddr),	city,	country)
	io.WriteString(w, "Lookup!")
} 

func getLookup(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.
	ipLookup()
}

func startServer(s http.Server, c context.CancelFunc) {
		err := s.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		c()
}

func generateServer(port int, mux *http.ServeMux, ctx context.Context) *http.Server{
	return &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/lookup", getLookup)
	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := generateServer(3333, mux, ctx,)
	//serverTwo := generateServer(44499, mux, ctx,)

	startServer(*serverOne, cancelCtx)
	//go startServer(*serverTwo, cancelCtx)

	<-ctx.Done()
}