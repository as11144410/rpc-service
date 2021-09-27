package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	helloworldApi "rpc-service/api/protobuf-spec/helloworldpb"
	"rpc-service/internal/app/service"
	"rpc-service/internal/pkg/database"
)

const port = ":8080"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	helloworldApi.RegisterGreeterServer(s, &service.HelloWorldServer{})

	reflection.Register(s)

	defer database.Db.Close()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
