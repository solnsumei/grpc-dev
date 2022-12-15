package main

import (
	"fmt"
	"log"
	"net"

	"github.com/solnsumei/grpc-dev/tutorial"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	tutorial.RegisterTutorialServer(s, &tutorial.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
