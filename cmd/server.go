package main

import (
	"log"
	"net"
	"os"

	"github.com/chickeniq/hypixel-go/internal/hypixel"
	"github.com/chickeniq/hypixel-go/internal/mojang"
	"google.golang.org/grpc"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is not set")
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	hypixel.Register(grpcServer, apiKey, nil)
	mojang.Register(grpcServer, nil)

	log.Fatal(grpcServer.Serve(listener))
}
