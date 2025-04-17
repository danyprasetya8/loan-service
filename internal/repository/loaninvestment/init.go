package loaninvestment

import (
	"loan-service/internal/entity"

	"gorm.io/gorm"
)

type ILoanInvestmentRepository interface {
	Create(en *entity.LoanInvestment) error
}

type LoanInvestment struct {
	db *gorm.DB
}

func New(db *gorm.DB) ILoanInvestmentRepository {
	return &LoanInvestment{db}
}

func (l *LoanInvestment) Create(en *entity.LoanInvestment) error {
	return l.db.Create(en).Error
}
