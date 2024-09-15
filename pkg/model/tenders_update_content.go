package model

import (
	"context"
	"tenders/pkg/api"

	"github.com/google/uuid"
)

func (m *Tenders) UpdateContent(ctx context.Context, tender *api.Tender) error {
	tx, err := m.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `
		INSERT INTO tender_content
			(name, description, type, version, tender_id)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING id
	`, tender.Name, tender.Description, tender.ServiceType, tender.Version, tender.ID)

	var nextContent uuid.UUID
	err = row.Scan(&nextContent)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
		UPDATE tender_content_ref
		SET content_id = $1
		WHERE tender_id = $2
	`, nextContent, tender.ID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
