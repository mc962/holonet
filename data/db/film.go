package db

import (
	"github.com/jackc/pgx"
	"holonet/data/resource"
)

type Film struct {
	DB *pgx.Conn
}

func FindFilm(id int) (resource.Film, error) {
	var film resource.Film
	row := DB.QueryRow("SELECT * FROM films WHERE id = $1;", id)

	err := row.Scan(
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

	if err == nil {
		return film, nil
	} else {
		return resource.Film{}, err
	}
}

func AllFilms() ([]resource.Film, error) {
	rows, err := DB.Query("SELECT * FROM films;")

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
