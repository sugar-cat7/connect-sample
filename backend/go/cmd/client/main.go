package main

import (
	"context"
	"log"
	"net/http"

	v1 "github.com/sugar-cat7/connect-sample/gen/proto/v1"             // generated by protoc-gen-go
	"github.com/sugar-cat7/connect-sample/gen/proto/v1/greetv1connect" // generated by protoc-gen-connect-go

	"connectrpc.com/connect"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		connect.WithGRPC(),
	)
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&v1.GreetRequest{Name: "Jane"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Greeting)
}