package handler

import (
	"context"
	"tenders/pkg/api"

	"github.com/google/uuid"
)

func (h *handler) GetUserTenders(ctx context.Context, params api.GetUserTendersParams) (api.GetUserTendersRes, error) {
	unauthorized := api.ErrorResponse(api.GetTenderStatusUnauthorized{})

	user, err := h.authUser(ctx, params.Username.Value)
	if err != nil {
		switch err {
		case ErrUnauthorized:
			return &unauthorized, nil
		default:
			return nil, err
		}
	}
	if user.Organization == uuid.Nil {
		return &unauthorized, nil
	}

	tenders, err := h.tenders.GetOfOrg(ctx,
		int(params.Limit.Value),
		int(params.Offset.Value),
		user.Organization)
	if err != nil {
		return nil, err
	}

	resp := api.GetUserTendersOKApplicationJSON(tenders)
	return &resp, nil
}
