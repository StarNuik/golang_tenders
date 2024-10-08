// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CheckServer implements checkServer operation.
//
// Этот эндпоинт используется для проверки готовности
// сервера обрабатывать запросы.
// Чекер программа будет ждать первый успешный ответ и
// затем начнет выполнение тестовых сценариев.
//
// GET /ping
func (UnimplementedHandler) CheckServer(ctx context.Context) (r CheckServerRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateBid implements createBid operation.
//
// Создание предложения для существующего тендера.
//
// POST /bids/new
func (UnimplementedHandler) CreateBid(ctx context.Context, req *CreateBidReq) (r CreateBidRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateTender implements createTender operation.
//
// Создание нового тендера с заданными параметрами.
//
// POST /tenders/new
func (UnimplementedHandler) CreateTender(ctx context.Context, req *CreateTenderReq) (r CreateTenderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// EditBid implements editBid operation.
//
// Редактирование существующего предложения.
//
// PATCH /bids/{bidId}/edit
func (UnimplementedHandler) EditBid(ctx context.Context, req *EditBidReq, params EditBidParams) (r EditBidRes, _ error) {
	return r, ht.ErrNotImplemented
}

// EditTender implements editTender operation.
//
// Изменение параметров существующего тендера.
//
// PATCH /tenders/{tenderId}/edit
func (UnimplementedHandler) EditTender(ctx context.Context, req *EditTenderReq, params EditTenderParams) (r EditTenderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBidReviews implements getBidReviews operation.
//
// Ответственный за организацию может посмотреть
// прошлые отзывы на предложения автора, который создал
// предложение для его тендера.
//
// GET /bids/{tenderId}/reviews
func (UnimplementedHandler) GetBidReviews(ctx context.Context, params GetBidReviewsParams) (r GetBidReviewsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBidStatus implements getBidStatus operation.
//
// Получить статус предложения по его уникальному
// идентификатору.
//
// GET /bids/{bidId}/status
func (UnimplementedHandler) GetBidStatus(ctx context.Context, params GetBidStatusParams) (r GetBidStatusRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBidsForTender implements getBidsForTender operation.
//
// Получение предложений, связанных с указанным
// тендером.
//
// GET /bids/{tenderId}/list
func (UnimplementedHandler) GetBidsForTender(ctx context.Context, params GetBidsForTenderParams) (r GetBidsForTenderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTenderStatus implements getTenderStatus operation.
//
// Получить статус тендера по его уникальному
// идентификатору.
//
// GET /tenders/{tenderId}/status
func (UnimplementedHandler) GetTenderStatus(ctx context.Context, params GetTenderStatusParams) (r GetTenderStatusRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTenders implements getTenders operation.
//
// Список тендеров с возможностью фильтрации по типу
// услуг.
// Если фильтры не заданы, возвращаются все тендеры.
//
// GET /tenders
func (UnimplementedHandler) GetTenders(ctx context.Context, params GetTendersParams) (r GetTendersRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserBids implements getUserBids operation.
//
// Получение списка предложений текущего пользователя.
// Для удобства использования включена поддержка
// пагинации.
//
// GET /bids/my
func (UnimplementedHandler) GetUserBids(ctx context.Context, params GetUserBidsParams) (r GetUserBidsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserTenders implements getUserTenders operation.
//
// Получение списка тендеров текущего пользователя.
// Для удобства использования включена поддержка
// пагинации.
//
// GET /tenders/my
func (UnimplementedHandler) GetUserTenders(ctx context.Context, params GetUserTendersParams) (r GetUserTendersRes, _ error) {
	return r, ht.ErrNotImplemented
}

// RollbackBid implements rollbackBid operation.
//
// Откатить параметры предложения к указанной версии.
// Это считается новой правкой, поэтому версия
// инкрементируется.
//
// PUT /bids/{bidId}/rollback/{version}
func (UnimplementedHandler) RollbackBid(ctx context.Context, params RollbackBidParams) (r RollbackBidRes, _ error) {
	return r, ht.ErrNotImplemented
}

// RollbackTender implements rollbackTender operation.
//
// Откатить параметры тендера к указанной версии. Это
// считается новой правкой, поэтому версия
// инкрементируется.
//
// PUT /tenders/{tenderId}/rollback/{version}
func (UnimplementedHandler) RollbackTender(ctx context.Context, params RollbackTenderParams) (r RollbackTenderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// SubmitBidDecision implements submitBidDecision operation.
//
// Отправить решение (одобрить или отклонить) по
// предложению.
//
// PUT /bids/{bidId}/submit_decision
func (UnimplementedHandler) SubmitBidDecision(ctx context.Context, params SubmitBidDecisionParams) (r SubmitBidDecisionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// SubmitBidFeedback implements submitBidFeedback operation.
//
// Отправить отзыв по предложению.
//
// PUT /bids/{bidId}/feedback
func (UnimplementedHandler) SubmitBidFeedback(ctx context.Context, params SubmitBidFeedbackParams) (r SubmitBidFeedbackRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateBidStatus implements updateBidStatus operation.
//
// Изменить статус предложения по его уникальному
// идентификатору.
//
// PUT /bids/{bidId}/status
func (UnimplementedHandler) UpdateBidStatus(ctx context.Context, params UpdateBidStatusParams) (r UpdateBidStatusRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateTenderStatus implements updateTenderStatus operation.
//
// Изменить статус тендера по его идентификатору.
//
// PUT /tenders/{tenderId}/status
func (UnimplementedHandler) UpdateTenderStatus(ctx context.Context, params UpdateTenderStatusParams) (r UpdateTenderStatusRes, _ error) {
	return r, ht.ErrNotImplemented
}
