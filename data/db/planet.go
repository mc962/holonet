package db

import (
	"github.com/jackc/pgx"
)

type Planet struct {
	DB *pgx.Conn
}
