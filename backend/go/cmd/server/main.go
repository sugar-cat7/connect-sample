package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/sugar-cat7/connect-sample/gen/proto/v1/greetv1connect" // generated by protoc-gen-connect-go
	"github.com/sugar-cat7/connect-sample/internal/infra/http/server"
)

func main() {
	greeter := server.NewGreetServiceHandler()
	mux := http.NewServeMux()

	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)
	log.Println("Starting server on localhost:8080")
	log.Fatal(http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	))
}