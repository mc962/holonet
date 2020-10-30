package db

import (
	"github.com/jackc/pgx"
)

type Species struct {
	DB *pgx.Conn
}
