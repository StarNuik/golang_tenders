package handler

import (
	"context"
	"tenders/pkg/api"
)

func (h *handler) EditTender(ctx context.Context, req *api.EditTenderReq, params api.EditTenderParams) (api.EditTenderRes, error) {
	tender, err := h.authTender(ctx, params.TenderId, params.Username)
	if err != nil {
		switch err {
		case ErrNotFound:
			return &api.EditTenderNotFound{}, nil
		case ErrUnauthorized:
			return &api.EditTenderUnauthorized{}, nil
		case ErrForbidden:
			return &api.EditTenderForbidden{}, nil
		default:
			return nil, err
		}
	}
	if !(req.Name.Set || req.Description.Set || req.ServiceType.Set) {
		return &api.EditTenderBadRequest{}, nil
	}

	tender.Name = req.Name.Or(tender.Name)
	tender.Description = req.Description.Or(tender.Description)
	tender.ServiceType = req.ServiceType.Or(tender.ServiceType)
	tender.Version += 1

	err = h.tenders.UpdateContent(ctx, tender)
	return tender, err
}
