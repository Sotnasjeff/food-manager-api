package main

import (
	"database/sql"
	"log"

	"github.com/Sotnasjeff/movies-manager-api/db"
	"github.com/Sotnasjeff/movies-manager-api/util"
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
	defer conn.Close()

	moviesDB := db.NewMovie(conn)
}
