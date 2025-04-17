package handler

import (
	"loan-service/internal/service/auth"
	"loan-service/internal/service/borrower"
	"loan-service/internal/service/loan"
)

type Handler struct {
	authService     auth.IAuthService
	borrowerService borrower.IBorrowerService
	loanService     loan.ILoanService
}

func New(
	authService auth.IAuthService,
	borrowerService borrower.IBorrowerService,
	loanService loan.ILoanService,
) *Handler {
	return &Handler{
		authService,
		borrowerService,
		loanService,
	}
}
