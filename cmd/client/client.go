package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	h "github.com/chickeniq/hypixel-go/proto/hypixel"
	m "github.com/chickeniq/hypixel-go/proto/mojang"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func run(ctx context.Context) error {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("localhost:50051", opts)
	if err != nil {
		return err
	}
	defer conn.Close()

	hypixel := h.NewHypixelClient(conn)
	mojang := m.NewMojangClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	profile, err := mojang.GetProfile(ctx, &m.ProfileRequest{Name: "ChickenIQ"})
	if err != nil {
		return err
	}

	response, err := hypixel.GetNetworth(ctx, &h.SkyBlockRequest{Uuid: profile.Id})
	if err != nil {
		return err
	}

	fmt.Printf("Total: %.2f\n", response.GetTotal())
	fmt.Printf("Non-cosmetic: %.2f\n", response.GetNonCosmetic())
	return nil
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
