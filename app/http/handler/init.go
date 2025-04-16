package handler

import (
	"loan-service/internal/service/auth"
	"loan-service/internal/service/borrower"
)

type Handler struct {
	authService     auth.IAuthService
	borrowerService borrower.IBorrowerService
}

func New(
	authService auth.IAuthService,
	borrowerService borrower.IBorrowerService,
) *Handler {
	return &Handler{authService, borrowerService}
}
