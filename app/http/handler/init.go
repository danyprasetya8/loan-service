package handler

import (
	"loan-service/internal/service/auth"
	"loan-service/internal/service/borrower"
	"loan-service/internal/service/file"
	"loan-service/internal/service/loan"
)

type Handler struct {
	fileService     file.IFileService
	authService     auth.IAuthService
	borrowerService borrower.IBorrowerService
	loanService     loan.ILoanService
}

func New(
	fileService file.IFileService,
	authService auth.IAuthService,
	borrowerService borrower.IBorrowerService,
	loanService loan.ILoanService,
) *Handler {
	return &Handler{
		fileService,
		authService,
		borrowerService,
		loanService,
	}
}
