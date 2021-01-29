package data

import (
	"context"
	"github.com/jackc/pgx/v4"
	"holonet/data/db"
	"log"
	"os"
)

func Initialize() *pgx.Conn {
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	db.DB = conn

	return conn
}
