package db

import (
	"github.com/jackc/pgx"
)

type Person struct {
	DB *pgx.Conn
}
