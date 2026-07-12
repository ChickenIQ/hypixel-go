package mojang

import (
	"context"

	"github.com/chickeniq/hypixel-go/pkg/mojang"
	pb "github.com/chickeniq/hypixel-go/proto/mojang"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedMojangServer
	client *mojang.Client
}

func Register(s grpc.ServiceRegistrar) {
	pb.RegisterMojangServer(s, &Server{client: mojang.NewClient()})
}

func (s *Server) GetProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileReponse, error) {
	p, err := s.client.GetProfile(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	return &pb.ProfileReponse{
		Name: p.Name,
		Id:   p.ID,
	}, nil
}
