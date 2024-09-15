package handler

import (
	"context"
	"tenders/pkg/api"

	"github.com/google/uuid"
)

func (h *handler) UpdateTenderStatus(ctx context.Context, params api.UpdateTenderStatusParams) (api.UpdateTenderStatusRes, error) {
	tender, err := h.authTender(ctx, params.TenderId, params.Username)
	if err != nil {
		switch err {
		case ErrNotFound:
			return &api.UpdateTenderStatusNotFound{}, nil
		case ErrUnauthorized:
			return &api.UpdateTenderStatusUnauthorized{}, nil
		case ErrForbidden:
			return &api.UpdateTenderStatusForbidden{}, nil
		default:
			return nil, err
		}
	}

	err = h.tenders.SetStatus(ctx, uuid.UUID(params.TenderId), params.Status)
	if err != nil {
		return nil, err
	}

	tender.Status = params.Status
	return tender, nil
}
