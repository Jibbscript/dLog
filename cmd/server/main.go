package main

import (
	"log"

	"github.com/jibbscript/dlog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
