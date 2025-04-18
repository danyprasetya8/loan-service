//go:generate mockgen -destination=../../../mocks/mock_loan_approval_repository.go -package=mocks loan-service/internal/repository/loanapproval ILoanApprovalRepository
package loanapproval

import (
	"loan-service/internal/entity"

	"gorm.io/gorm"
)

type ILoanApprovalRepository interface {
	Create(en *entity.LoanApproval) error
}

type LoanApproval struct {
	db *gorm.DB
}

func New(db *gorm.DB) ILoanApprovalRepository {
	return &LoanApproval{db}
}

func (l *LoanApproval) Create(en *entity.LoanApproval) error {
	return l.db.Create(en).Error
}
