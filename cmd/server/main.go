package main

import (
	"log"

	httpserver "github.com/jibbscript/dlog/internal/server/httpPrototype"
)

func main() {
	srv := httpserver.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
