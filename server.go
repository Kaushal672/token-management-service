package main

import (
	"log"
	"net"
	"token-management-service/grpcHandler"
	"token-management-service/protogen/token"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Failed to listen", err)
	}

	grpcServer := grpc.NewServer()
	tokenServer := grpcHandler.NewTokenServer()
	token.RegisterTokenServer(grpcServer, &tokenServer)

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to server", err)
	}
}
