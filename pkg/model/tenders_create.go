package model

import (
	"context"
	"tenders/pkg/api"

	"github.com/google/uuid"
)

func (m *Tenders) Create(ctx context.Context, req *api.CreateTenderReq) (*api.Tender, error) {
	tender := api.Tender{
		Status:         req.Status,
		OrganizationId: req.OrganizationId,
		Name:           req.Name,
		Description:    req.Description,
		ServiceType:    req.ServiceType,
	}

	tx, err := m.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `
		INSERT INTO tender
			(status, organization_id)
		VALUES
			($1, $2)
		RETURNING
			id, created_at
	`, req.Status, req.OrganizationId)

	err = row.Scan(&tender.ID, &tender.CreatedAt)
	if err != nil {
		return nil, err
	}

	row = tx.QueryRow(ctx, `
		INSERT INTO tender_content
			(name, description, type, tender_id)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			id, version
	`, req.Name, req.Description, req.ServiceType, tender.ID)

	var contentId uuid.UUID
	err = row.Scan(&contentId, &tender.Version)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO tender_content_ref
			(tender_id, content_id)
		VALUES
			($1, $2)
	`, tender.ID, contentId)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &tender, nil
}
