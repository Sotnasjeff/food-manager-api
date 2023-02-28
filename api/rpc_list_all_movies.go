package api

import (
	"context"

	"github.com/Sotnasjeff/movies-manager-api/pb"
)

func (server *Server) ListAllMovies(ctx context.Context, p *pb.Blank) (*pb.ListMoviesResponse, error) {
	movies, err := server.store.FindAllMovies()
	if err != nil {
		return nil, err
	}

	var moviesResponse []*pb.Movie

	for _, movie := range movies {
		movieResp := &pb.Movie{
			Id:          movie.ID,
			Name:        movie.Name,
			Description: movie.Description,
		}
		moviesResponse = append(moviesResponse, movieResp)
	}

	return &pb.ListMoviesResponse{
		Movie: moviesResponse,
	}, nil
}
