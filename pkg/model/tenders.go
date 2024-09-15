package model

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Tenders struct {
	db *pgxpool.Pool
}

func NewTendersModel(db *pgxpool.Pool) *Tenders {
	return &Tenders{
		db: db,
	}
}
