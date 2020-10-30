package db

import "database/sql"

type Person struct {
	DB *sql.DB
}
