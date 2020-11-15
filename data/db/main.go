package db

import (
	"github.com/jackc/pgx"
)

// TODO this isn't a great design, but doing it here for easy app prototyping
var DB *pgx.Conn
