package service

import (
	"context"
	"log"
	hw "rpc-service/api/protobuf-spec/helloworldpb"
)

type HelloWorldServer struct {
	hw.UnimplementedGreeterServer
}

func (s *HelloWorldServer) SayHello(ctx context.Context, in *hw.HelloRequest) (*hw.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hw.HelloReply{Message: "Hello " + in.GetName()}, nil
}
