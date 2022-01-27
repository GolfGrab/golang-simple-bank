package main

import (
	"database/sql"
	"log"

	"github.com/GolfGrab/journey-to-complete-backend/api"
	db "github.com/GolfGrab/journey-to-complete-backend/db/sqlc"
	"github.com/GolfGrab/journey-to-complete-backend/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
