package api

import (
	"context"
	"io"

	"github.com/Sotnasjeff/movies-manager-api/pb"
)

func (server *Server) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	arg, err := server.store.CreateMovie(req.GetName(), req.GetDescription())
	if err != nil {
		return nil, err
	}

	return &pb.CreateMovieResponse{
		Movie: &pb.Movie{
			Id:          arg.ID,
			Name:        arg.Name,
			Description: arg.Description,
		},
	}, nil
}

func (server *Server) CreateMovieStream(stream pb.MovieService_CreateMovieStreamServer) error {
	movies := &pb.ListMoviesResponse{}

	for {
		movie, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(movies)
		}
		if err != nil {
			return err
		}

		moviesResult, err := server.store.CreateMovie(movie.GetName(), movie.GetDescription())
		if err != nil {
			return err
		}

		movies.Movie = append(movies.Movie, &pb.Movie{
			Id:          moviesResult.ID,
			Name:        moviesResult.Name,
			Description: moviesResult.Description,
		})
	}
}

func (server *Server) CreateMovieBidirectionalStream(stream pb.MovieService_CreateMovieBidirectionalStreamServer) error {
	for {
		movie, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		moviesResult, err := server.store.CreateMovie(movie.GetName(), movie.GetDescription())
		if err != nil {
			return err
		}

		err = stream.Send(&pb.CreateMovieResponse{
			Movie: &pb.Movie{
				Id:          moviesResult.ID,
				Name:        moviesResult.Name,
				Description: moviesResult.Description,
			},
		})

		if err != nil {
			return err
		}
	}
}
