//go:generate mockgen -destination=../../../mocks/mock_loan_disbursement_repository.go -package=mocks loan-service/internal/repository/loandisbursement ILoanDisbursementRepository
package loandisbursement

import (
	"loan-service/internal/entity"

	"gorm.io/gorm"
)

type ILoanDisbursementRepository interface {
	Create(en *entity.LoanDisbursement) error
}

type LoanDisbursement struct {
	db *gorm.DB
}

func New(db *gorm.DB) ILoanDisbursementRepository {
	return &LoanDisbursement{db}
}

func (l *LoanDisbursement) Create(en *entity.LoanDisbursement) error {
	return l.db.Create(en).Error
}
