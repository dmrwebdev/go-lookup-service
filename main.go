package main

import (
	"golookupservice/pkg/server"
)

func main() {
	s := server.NewServer(3333)
	server.StartServer(s)
}