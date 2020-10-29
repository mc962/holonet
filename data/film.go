package data

import "database/sql"

type Film struct {
	Title        string     `json:"title"`
	EpisodeId    int        `json:"episode_id"`
	OpeningCrawl string     `json:"opening_crawl"`
	Director     string     `json:"director"`
	Producer     string     `json:"producer"`
	ReleaseDate  string     `json:"release_date"`
	Characters   []Person   `json:"characters"`
	Planets      []Planet   `json:"planets"`
	Starships    []Starship `json:"starships"`
	Vehicles     []Vehicle  `json:"vehicles"`
	Species      []Species  `json:"species"`
	Metadata
	db *sql.DB
}

func (film Film) All(db *sql.DB) ([]Film, error) {
	rows, err := db.Query("SELECT * FROM films;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var films []Film

	for rows.Next() {
		var film Film

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
