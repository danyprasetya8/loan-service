package handler

import "loan-service/internal/service/borrower"

type Handler struct {
	borrowerService borrower.IBorrowerService
}

func New(
	borrowerService borrower.IBorrowerService,
) *Handler {
	return &Handler{borrowerService}
}
