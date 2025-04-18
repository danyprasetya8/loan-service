//go:generate mockgen -destination=../../../mocks/mock_loan_investment_repository.go -package=mocks loan-service/internal/repository/loaninvestment ILoanInvestmentRepository
package loaninvestment

import (
	"loan-service/internal/entity"

	"gorm.io/gorm"
)

type ILoanInvestmentRepository interface {
	GetByLoanID(loanID string) []entity.LoanInvestment
	GetByLoanAndInvestor(loanID, investorID string) *entity.LoanInvestment
	Create(en *entity.LoanInvestment) error
	Save(en *entity.LoanInvestment) error
}

type LoanInvestment struct {
	db *gorm.DB
}

func New(db *gorm.DB) ILoanInvestmentRepository {
	return &LoanInvestment{db}
}

func (l *LoanInvestment) GetByLoanID(loanID string) []entity.LoanInvestment {
	li := make([]entity.LoanInvestment, 0)

	l.db.Where("loan_id = ?", loanID).
		Find(&li)
	return li
}

func (l *LoanInvestment) GetByLoanAndInvestor(loanID, investorID string) *entity.LoanInvestment {
	loanInvestment := &entity.LoanInvestment{}

	err := l.db.Where("loan_id = ?", loanID).
		Where("investor_id = ?", investorID).
		First(&loanInvestment).
		Error

	if err != nil {
		return nil
	}

	return loanInvestment
}

func (l *LoanInvestment) Create(en *entity.LoanInvestment) error {
	return l.db.Create(en).Error
}

func (l *LoanInvestment) Save(en *entity.LoanInvestment) error {
	return l.db.Save(en).Error
}
