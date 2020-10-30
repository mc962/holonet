package db

import (
	"database/sql"
	"holonet/data/resource"
)

type Film struct {
	DB *sql.DB
}

func (film Film) All(db *sql.DB) ([]resource.Film, error) {
	rows, err := db.Query("SELECT * FROM films;")

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
