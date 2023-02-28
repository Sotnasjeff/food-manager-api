package api

import (
	"github.com/Sotnasjeff/movies-manager-api/db"
	"github.com/Sotnasjeff/movies-manager-api/pb"
)

type Server struct {
	pb.UnimplementedMovieServiceServer
	store db.Movie
}

func NewServer(store db.Movie) (*Server, error) {
	return &Server{
		store: store,
	}, nil
}
