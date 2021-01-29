package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"holonet/data/resource"
)

type Species struct {
	DB *pgx.Conn
}

func FindSpecies(id int) (resource.Species, error) {
	var species resource.Species
	row := DB.QueryRow(context.Background(), "SELECT * FROM species WHERE id = $1;", id)

	err := row.Scan(
		&species.Name,
		&species.Classification,
		&species.Designation,
		&species.AverageHeight,
		&species.AverageLifespan,
		&species.HairColors,
		&species.SkinColors,
		&species.EyeColors,
		&species.Homeworld,
		&species.Language,
		&species.People,
		&species.Films,
	)

	if err == nil {
		return species, nil
	} else {
		return resource.Species{}, err
	}
}

func AllSpecies() ([]resource.Species, error) {
	rows, err := DB.Query(context.Background(), "SELECT * FROM species;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allSpecies []resource.Species

	for rows.Next() {
		var species resource.Species

		err := rows.Scan(
			&species.Name,
			&species.Classification,
			&species.Designation,
			&species.AverageHeight,
			&species.AverageLifespan,
			&species.HairColors,
			&species.SkinColors,
			&species.EyeColors,
			&species.Homeworld,
			&species.Language,
			&species.People,
			&species.Films,
		)
		if err != nil {
			return nil, err
		}

		allSpecies = append(allSpecies, species)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return allSpecies, nil
}
