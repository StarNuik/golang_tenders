package model

import (
	"context"
	"tenders/pkg/api"
)

type GetTendersParams struct {
	Limit       int
	Offset      int
	TypeFilters []string
}

func (m *Tenders) GetPublished(ctx context.Context, limit int, offset int, filters []string) ([]api.Tender, error) {
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
		WHERE status = 'Published' AND type = ANY($1)
		ORDER BY name ASC
		OFFSET $2
		LIMIT $3
	`, filters, offset, limit)
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
