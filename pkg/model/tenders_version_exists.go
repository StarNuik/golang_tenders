package model

import (
	"context"

	"github.com/google/uuid"
)

func (m *Tenders) VersionExists(ctx context.Context, tender uuid.UUID, version int) (bool, error) {
	row := m.db.QueryRow(ctx, `
		SELECT COUNT(id)
		FROM tender_content
		WHERE tender_id = $1 AND version = $2
	`, tender, version)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count >= 1, nil
}
