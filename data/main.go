package data

import (
	"database/sql"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"holonet/data/db"
	"log"
)

func Initialize() Models {
	database, err := sql.Open("pgx", "postgres://holonet_local:local@localhost:5433/holonet_dev")
	if err != nil {
		log.Fatalln(err)
	}
	models := RegisterModels(database)

	return models
}

func RegisterModels(database *sql.DB) Models {
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
