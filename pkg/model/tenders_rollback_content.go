package model

import (
	"context"
	"tenders/pkg/api"

	"github.com/google/uuid"
)

func (m *Tenders) RollbackContent(ctx context.Context, tender *api.Tender, targetVersion int) (*api.Tender, error) {
	tx, err := m.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `
		INSERT INTO tender_content
			(tender_id, name, description, type, version)
			SELECT $1, name, description, type, $2
			FROM tender_content
			WHERE tender_id = $1 AND version = $3
		RETURNING id, name, description, type, version`,
		tender.ID, tender.Version+1, targetVersion)

	var contentId uuid.UUID
	err = row.Scan(&contentId, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Version)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, `
		UPDATE tender_content_ref
		SET content_id = $1
		WHERE tender_id = $2
	`, contentId, tender.ID)
	if err != nil {
		return nil, err
	}

	return tender, tx.Commit(ctx)
}
