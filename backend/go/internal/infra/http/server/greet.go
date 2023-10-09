package server

import (
	"context"
	"fmt"
	"log"

	"connectrpc.com/connect"
	v1 "github.com/sugar-cat7/connect-sample/gen/proto/v1"
)

type GreetServiceHandlerImpl struct {
}

func NewGreetServiceHandler() *GreetServiceHandlerImpl {
	return &GreetServiceHandlerImpl{}
}

func (s *GreetServiceHandlerImpl) Greet(ctx context.Context, req *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&v1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
