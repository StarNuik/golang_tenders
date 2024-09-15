package handler

import (
	"context"
	"errors"
	"tenders/pkg/api"

	"github.com/google/uuid"
)

func (h *handler) GetTenderStatus(ctx context.Context, params api.GetTenderStatusParams) (api.GetTenderStatusRes, error) {
	tender, err := h.tenders.Find(ctx, uuid.UUID(params.TenderId))
	if err != nil {
		return nil, err
	}
	if tender == nil {
		return &api.GetTenderStatusNotFound{}, nil
	}

	if tender.Status == api.TenderStatusPublished {
		return &tender.Status, nil
	}
	if !params.Username.Set {
		return &api.GetTenderStatusForbidden{}, nil
	}

	user, err := h.authUser(ctx, params.Username.Value)
	if errors.Is(err, ErrUnauthorized) {
		return &api.GetTenderStatusUnauthorized{}, nil
	}
	if err != nil {
		return nil, err
	}

	if user.Organization != uuid.UUID(tender.OrganizationId) {
		return &api.GetTenderStatusForbidden{}, nil
	}

	return &tender.Status, nil
}
