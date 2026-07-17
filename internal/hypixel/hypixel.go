package hypixel

import (
	"github.com/chickeniq/hypixel-go/pkg/cache"
	"github.com/chickeniq/hypixel-go/pkg/hypixel"
	pb "github.com/chickeniq/hypixel-go/proto/hypixel"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHypixelServer
	client *hypixel.Client
}

func Register(s grpc.ServiceRegistrar, apiKey string, cache *cache.Cache) {
	pb.RegisterHypixelServer(s, &Server{client: hypixel.NewClient(apiKey, cache)})
}
