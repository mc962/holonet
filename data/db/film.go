package db

import (
	"github.com/jackc/pgx"
	"holonet/data/resource"
)

type Film struct {
	DB *pgx.Conn
}

func (film Film) All() ([]resource.Film, error) {
	rows, err := film.DB.Query("SELECT * FROM films;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var films []resource.Film

	for rows.Next() {
		var film resource.Film

		err := rows.Scan(
			&film.Title,
			&film.EpisodeId,
			&film.OpeningCrawl,
			&film.Director,
			&film.Producer,
			&film.ReleaseDate,
			&film.Characters,
			&film.Planets,
			&film.Starships,
			&film.Vehicles,
			&film.Species,
		)
		if err != nil {
			return nil, err
		}

		films = append(films, film)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return films, nil
}
