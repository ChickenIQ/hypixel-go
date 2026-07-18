package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/chickeniq/hypixel-go/internal/hypixel"
	"github.com/chickeniq/hypixel-go/internal/mojang"
	"google.golang.org/grpc"
)

func run(ctx context.Context) error {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return errors.New("API_KEY is not set")
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	hypixel.Register(grpcServer, apiKey, nil)
	mojang.Register(grpcServer, nil)

	go func() {
		<-ctx.Done()
		grpcServer.GracefulStop()
	}()

	return grpcServer.Serve(listener)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		if ctx.Err() != nil {
			return
		}
		log.Fatal(err)
	}
}
