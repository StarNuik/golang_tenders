package model

import (
	"context"
	"errors"
	"tenders/pkg/api"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (m *Tenders) Find(ctx context.Context, tender uuid.UUID) (*api.Tender, error) {
	row := m.db.QueryRow(ctx, `
		SELECT
			tender.id,
			name,
			description,
			type,
			status,
			organization_id,
			version,
			created_at
		FROM tender_content_ref ref
		JOIN tender
		ON ref.tender_id = $1 AND tender.id = ref.tender_id
		JOIN tender_content
		ON ref.content_id = tender_content.id
	`, tender)

	out := api.Tender{}
	err := row.Scan(
		&out.ID,
		&out.Name,
		&out.Description,
		&out.ServiceType,
		&out.Status,
		&out.OrganizationId,
		&out.Version,
		&out.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &out, nil
}
