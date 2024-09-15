package model

import (
	"context"
	"tenders/pkg/api"

	"github.com/google/uuid"
)

func (m *Tenders) SetStatus(ctx context.Context, tender uuid.UUID, status api.TenderStatus) error {
	_, err := m.db.Exec(ctx, `
		UPDATE tender
		SET status = $1
		WHERE id = $2`,
		status, tender)
	return err
}
