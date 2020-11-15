package db

import (
	"github.com/jackc/pgx"
	"holonet/data/resource"
)

type Planet struct {
	DB *pgx.Conn
}

func FindPlanet(id int) (resource.Planet, error) {
	var planet resource.Planet
	row := DB.QueryRow("SELECT * FROM planets WHERE id = $1;", id)

	err := row.Scan(
		&planet.Name,
		&planet.RotationPeriod,
		&planet.OrbitalPeriod,
		&planet.Diameter,
		&planet.Climate,
		&planet.Gravity,
		&planet.Terrain,
		&planet.SurfaceWater,
		&planet.Population,
		&planet.Residents,
		&planet.Films,
	)

	if err == nil {
		return planet, nil
	} else {
		return resource.Planet{}, err
	}
}

func AllPlanets() ([]resource.Planet, error) {
	rows, err := DB.Query("SELECT * FROM planets;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var planets []resource.Planet

	for rows.Next() {
		var planet resource.Planet

		err := rows.Scan(
			&planet.Name,
			&planet.RotationPeriod,
			&planet.OrbitalPeriod,
			&planet.Diameter,
			&planet.Climate,
			&planet.Gravity,
			&planet.Terrain,
			&planet.SurfaceWater,
			&planet.Population,
			&planet.Residents,
			&planet.Films,
		)
		if err != nil {
			return nil, err
		}

		planets = append(planets, planet)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return planets, nil
}
