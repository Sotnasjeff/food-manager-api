package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/Sotnasjeff/movies-manager-api/api"
	"github.com/Sotnasjeff/movies-manager-api/db"
	"github.com/Sotnasjeff/movies-manager-api/pb"
	"github.com/Sotnasjeff/movies-manager-api/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config:")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Could not connect to db:")
	}

	movieDB := db.NewMovie(conn)

	runGRPCServer(config, *movieDB)
}

func runGRPCServer(config util.Config, store db.Movie) {
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("Could not connect to server:")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMovieServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Could not connect to the port:")
	}

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatal("Could not connect to grpc Server:")
	}
}
