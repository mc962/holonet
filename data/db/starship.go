package db

import (
	"github.com/jackc/pgx"
	"holonet/data/resource"
)

type Starship struct {
	DB *pgx.Conn
}

func FindStarship(id int) (resource.Starship, error) {
	var starship resource.Starship
	row := DB.QueryRow("SELECT * FROM starships WHERE id = $1;", id)

	err := row.Scan(
		&starship.Name,
		&starship.Model,
		&starship.Manufacturer,
		&starship.CostInCredits,
		&starship.Length,
		&starship.MaxAtmospheringSpeed,
		&starship.Crew,
		&starship.Passengers,
		&starship.CargoCapacity,
		&starship.Consumables,
		&starship.HyperdriveRating,
		&starship.MGLT,
		&starship.StarshipClass,
		&starship.Pilots,
		&starship.Films,
	)

	if err == nil {
		return starship, nil
	} else {
		return resource.Starship{}, err
	}
}

func AllStarships() ([]resource.Starship, error) {
	rows, err := DB.Query("SELECT * FROM starships;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var starships []resource.Starship

	for rows.Next() {
		var starship resource.Starship

		err := rows.Scan(
			&starship.Name,
			&starship.Model,
			&starship.Manufacturer,
			&starship.CostInCredits,
			&starship.Length,
			&starship.MaxAtmospheringSpeed,
			&starship.Crew,
			&starship.Passengers,
			&starship.CargoCapacity,
			&starship.Consumables,
			&starship.HyperdriveRating,
			&starship.MGLT,
			&starship.StarshipClass,
			&starship.Pilots,
			&starship.Films,
		)
		if err != nil {
			return nil, err
		}

		starships = append(starships, starship)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return starships, nil
}
