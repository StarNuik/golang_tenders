package handler

import (
	"tenders/pkg/api"
	"tenders/pkg/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type handler struct {
	tenders *model.Tenders
	users   *model.Users
	db      *pgxpool.Pool
}

func New(db *pgxpool.Pool) api.Handler {
	h := handler{
		tenders: model.NewTendersModel(db),
		users:   model.NewUsersModel(db),
		db:      db,
	}
	return &h
}
