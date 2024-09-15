package handler

import (
	"context"
	"fmt"
	"tenders/pkg/api"
)

func (sv *handler) CreateBid(ctx context.Context, req *api.CreateBidReq) (api.CreateBidRes, error) {
	return nil, fmt.Errorf("not implemented")
}

func (sv *handler) SubmitBidDecision(ctx context.Context, params api.SubmitBidDecisionParams) (api.SubmitBidDecisionRes, error) {
	return nil, fmt.Errorf("not implemented")
}

func (sv *handler) GetBidsForTender(ctx context.Context, params api.GetBidsForTenderParams) (api.GetBidsForTenderRes, error) {
	return nil, fmt.Errorf("not implemented")
}
func (sv *handler) GetUserBids(ctx context.Context, params api.GetUserBidsParams) (api.GetUserBidsRes, error) {
	return nil, fmt.Errorf("not implemented")
}

func (sv *handler) GetBidStatus(ctx context.Context, params api.GetBidStatusParams) (api.GetBidStatusRes, error) {
	return nil, fmt.Errorf("not implemented")
}
func (sv *handler) UpdateBidStatus(ctx context.Context, params api.UpdateBidStatusParams) (api.UpdateBidStatusRes, error) {
	return nil, fmt.Errorf("not implemented")
}

func (sv *handler) EditBid(ctx context.Context, req *api.EditBidReq, params api.EditBidParams) (api.EditBidRes, error) {
	return nil, fmt.Errorf("not implemented")
}
func (sv *handler) RollbackBid(ctx context.Context, params api.RollbackBidParams) (api.RollbackBidRes, error) {
	return nil, fmt.Errorf("not implemented")
}

func (sv *handler) SubmitBidFeedback(ctx context.Context, params api.SubmitBidFeedbackParams) (api.SubmitBidFeedbackRes, error) {
	return nil, fmt.Errorf("not implemented")
}
func (sv *handler) GetBidReviews(ctx context.Context, params api.GetBidReviewsParams) (api.GetBidReviewsRes, error) {
	return nil, fmt.Errorf("not implemented")
}
