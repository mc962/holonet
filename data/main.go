package data

import (
	"github.com/jackc/pgx"
	"holonet/data/db"
	"log"
)

func Initialize() (*pgx.Conn, Models) {
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

	models := RegisterModels(conn)

	return conn, models
}

func RegisterModels(database *pgx.Conn) Models {
	return Models{
		Film:     db.Film{DB: database},
		Person:   db.Person{DB: database},
		Planet:   db.Planet{DB: database},
		Species:  db.Species{DB: database},
		Starship: db.Starship{DB: database},
		Vehicle:  db.Vehicle{DB: database},
	}
}

type Models struct {
	Film     db.Film
	Person   db.Person
	Planet   db.Planet
	Species  db.Species
	Starship db.Starship
	Vehicle  db.Vehicle
}
