package handler

import (
	"context"
	"strings"
	"tenders/pkg/api"
)

func (h *handler) CheckServer(ctx context.Context) (api.CheckServerRes, error) {
	return &api.CheckServerOK{
		Data: strings.NewReader("ok"),
	}, nil
}
