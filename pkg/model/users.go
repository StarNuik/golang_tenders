package model

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Users struct {
	db *pgxpool.Pool
}

type User struct {
	Id           uuid.UUID
	Username     string
	Organization uuid.UUID
}

func NewUsersModel(db *pgxpool.Pool) *Users {
	return &Users{
		db: db,
	}
}

func (m *Users) Find(ctx context.Context, username string) (*User, error) {
	user := User{
		Username: username,
	}

	row := m.db.QueryRow(ctx, `
		SELECT
			employee.id,
			ref.organization_id
		FROM employee
		LEFT JOIN organization_responsible ref
		ON employee.id = ref.user_id
		WHERE username = $1
	`, username)

	err := row.Scan(&user.Id, &user.Organization)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
