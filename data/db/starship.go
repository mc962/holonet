package db

import (
	"github.com/jackc/pgx"
)

type Starship struct {
	DB *pgx.Conn
}
