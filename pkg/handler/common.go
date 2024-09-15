package handler

import (
	"context"
	"errors"
	"tenders/pkg/api"
	"tenders/pkg/model"

	"github.com/google/uuid"
)

var (
	ErrNotFound     = errors.New("404")
	ErrUnauthorized = errors.New("401")
	ErrForbidden    = errors.New("403")
)

func (h *handler) authUser(ctx context.Context, username api.Username) (*model.User, error) {
	if len(username) == 0 {
		return nil, ErrUnauthorized
	}
	user, err := h.users.Find(ctx, string(username))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUnauthorized
	}
	return user, nil
}

func (h *handler) authTender(
	ctx context.Context,
	tenderId api.TenderId,
	username api.Username,
) (*api.Tender, error) {
	tender, err := h.tenders.Find(ctx, uuid.UUID(tenderId))
	if err != nil {
		return nil, err
	}
	if tender == nil {
		return nil, ErrNotFound
	}

	if len(username) == 0 {
		return nil, ErrUnauthorized
	}

	user, err := h.users.Find(ctx, string(username))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUnauthorized
	}
	if user.Organization != uuid.UUID(tender.OrganizationId) {
		return nil, ErrForbidden
	}

	return tender, nil
}
