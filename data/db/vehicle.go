package db

import (
	"github.com/jackc/pgx"
)

type Vehicle struct {
	DB *pgx.Conn
}
