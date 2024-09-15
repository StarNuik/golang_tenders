package model

import (
	"context"
	"tenders/pkg/api"

	"github.com/google/uuid"
)

func (m *Tenders) GetOfOrg(ctx context.Context, limit int, offset int, org uuid.UUID) ([]api.Tender, error) {
	rows, err := m.db.Query(ctx, `
		SELECT
			tender.id,
			name,
			description,
			type,
			status,
			organization_id,
			version,
			created_at
		FROM tender_content_ref
		JOIN tender
		ON tender.id = tender_content_ref.tender_id
		JOIN tender_content
		ON tender_content.id = tender_content_ref.content_id
		WHERE organization_id = $1
		ORDER BY name ASC
		OFFSET $2
		LIMIT $3
	`, org, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []api.Tender{}
	for rows.Next() {
		tender := api.Tender{}
		err := rows.Scan(
			&tender.ID,
			&tender.Name,
			&tender.Description,
			&tender.ServiceType,
			&tender.Status,
			&tender.OrganizationId,
			&tender.Version,
			&tender.CreatedAt)
		if err != nil {
			return nil, err
		}

		out = append(out, tender)
	}

	return out, nil
}
