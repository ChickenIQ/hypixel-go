package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chickeniq/hypixel-go/internal/hypixel"
	"github.com/chickeniq/hypixel-go/internal/mojang"
	"github.com/chickeniq/hypixel-go/pkg/cache"
	"google.golang.org/grpc"
)

func run(ctx context.Context) error {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return fmt.Errorf("API_KEY is not set")
	}

	valkeyURL := os.Getenv("VALKEY_URL")
	var valkeyCache *cache.Cache
	if valkeyURL != "" {
		var err error
		valkeyCache, err = cache.NewCache(valkeyURL, 15*time.Minute, 30*time.Second)
		if err != nil {
			return fmt.Errorf("create valkey cache: %w", err)
		}
		defer valkeyCache.Close()
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	hypixel.Register(grpcServer, apiKey, valkeyCache)
	mojang.Register(grpcServer, valkeyCache)

	go func() {
		<-ctx.Done()
		grpcServer.GracefulStop()
	}()

	return grpcServer.Serve(listener)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}
