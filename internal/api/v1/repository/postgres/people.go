package postgres

import "github.com/jmoiron/sqlx"

type People struct {
	db *sqlx.DB
}

func NewPeople(db *sqlx.DB) *People {
	return &People{
		db: db,
	}
}
