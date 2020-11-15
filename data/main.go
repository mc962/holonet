package data

import (
	"github.com/jackc/pgx"
	"holonet/data/db"
	"log"
)

func Initialize() *pgx.Conn {
	config := pgx.ConnConfig{
		Host:     "localhost",
		Port:     5433,
		Database: "holonet_dev",
		User:     "holonet_local",
		Password: "local",
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	db.DB = conn

	return conn
}
