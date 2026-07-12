package hypixel

import (
	"context"

	pb "github.com/chickeniq/hypixel-go/proto/hypixel"
)

func (s *Server) GetNetworth(ctx context.Context, req *pb.SkyBlockRequest) (*pb.NetworthResponse, error) {
	nw, err := s.client.GetNetworth(ctx, req.GetUuid(), req.GetProfileName())
	if err != nil {
		return nil, err
	}

	return &pb.NetworthResponse{
		Total:       nw.Total,
		NonCosmetic: nw.NonCosmetic,
	}, nil
}
