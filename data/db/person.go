package db

import (
	"github.com/jackc/pgx"
	"holonet/data/resource"
)

type Person struct {
	DB *pgx.Conn
}

func FindPerson(id int) (resource.Person, error) {
	var person resource.Person
	row := DB.QueryRow("SELECT * FROM people WHERE id = $1;", id)

	err := row.Scan(
		&person.Name,
		&person.Height,
		&person.Mass,
		&person.HairColor,
		&person.SkinColor,
		&person.EyeColor,
		&person.BirthYear,
		&person.Gender,
		&person.Homeworld,
		&person.Films,
		&person.Species,
		&person.Vehicles,
		&person.Starships,
	)

	if err == nil {
		return person, nil
	} else {
		return resource.Person{}, err
	}
}

func AllPeople() ([]resource.Person, error) {
	rows, err := DB.Query("SELECT * FROM people;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []resource.Person

	for rows.Next() {
		var person resource.Person

		err := rows.Scan(
			&person.Name,
			&person.Height,
			&person.Mass,
			&person.HairColor,
			&person.SkinColor,
			&person.EyeColor,
			&person.BirthYear,
			&person.Gender,
			&person.Homeworld,
			&person.Films,
			&person.Species,
			&person.Vehicles,
			&person.Starships,
		)
		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return people, nil
}
