package handler

import (
	"context"
	"tenders/pkg/api"

	"github.com/google/uuid"
)

func (h *handler) CreateTender(ctx context.Context, req *api.CreateTenderReq) (api.CreateTenderRes, error) {
	user, err := h.authUser(ctx, req.CreatorUsername)
	if err != nil {
		switch err {
		case ErrUnauthorized:
			return &api.CreateTenderUnauthorized{}, nil
		default:
			return nil, err
		}
	}
	if user.Organization == uuid.Nil {
		return &api.CreateTenderForbidden{}, nil
	}

	if user.Organization != uuid.UUID(req.OrganizationId) {
		return &api.CreateTenderForbidden{}, nil
	}

	tender, err := h.tenders.Create(ctx, req)

	return tender, err
}
