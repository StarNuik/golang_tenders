package handler

import (
	"context"
	"tenders/pkg/api"
)

func (h *handler) GetTenders(ctx context.Context, params api.GetTendersParams) (api.GetTendersRes, error) {
	types := params.ServiceType
	if len(types) == 0 {
		types = api.TenderServiceTypeConstruction.AllValues()
	}

	filters := []string{}
	for _, tenderType := range types {
		filters = append(filters, string(tenderType))
	}

	tenders, err := h.tenders.GetPublished(ctx,
		int(params.Limit.Value),
		int(params.Offset.Value),
		filters)
	if err != nil {
		return nil, err
	}

	resp := api.GetTendersOKApplicationJSON(tenders)
	return &resp, nil
}
