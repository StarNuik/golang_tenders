package handler

import (
	"context"
	"tenders/pkg/api"

	"github.com/google/uuid"
)

func (h *handler) RollbackTender(ctx context.Context, params api.RollbackTenderParams) (api.RollbackTenderRes, error) {
	tender, err := h.authTender(ctx, params.TenderId, params.Username)
	if err != nil {
		switch err {
		case ErrNotFound:
			return &api.RollbackTenderNotFound{}, nil
		case ErrUnauthorized:
			return &api.RollbackTenderUnauthorized{}, nil
		case ErrForbidden:
			return &api.RollbackTenderForbidden{}, nil
		default:
			return nil, err
		}
	}
	id, targetVersion := uuid.UUID(params.TenderId), int(params.Version)
	ok, err := h.tenders.VersionExists(ctx, id, targetVersion)
	if err != nil {
		return nil, err
	}
	if !ok {
		return &api.RollbackTenderNotFound{}, nil
	}

	res, err := h.tenders.RollbackContent(ctx, tender, targetVersion)
	if err != nil {
		return nil, err
	}

	return res, err
}
