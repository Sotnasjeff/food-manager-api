package api

import (
	"context"

	"github.com/Sotnasjeff/movies-manager-api/pb"
)

func (server *Server) GetMovieById(ctx context.Context, req *pb.GetMovieByIdRequest) (*pb.GetMovieByIdResponse, error) {
	arg, err := server.store.FindMovieById(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetMovieByIdResponse{
		Movie: &pb.Movie{
			Id:          arg.ID,
			Name:        arg.Name,
			Description: arg.Description,
		},
	}, nil
}
