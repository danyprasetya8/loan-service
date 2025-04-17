package loan

import (
	"loan-service/internal/entity"

	"gorm.io/gorm"
)

type ILoanRepository interface {
	Create(en *entity.Loan) error
}

type Loan struct {
	db *gorm.DB
}

func New(db *gorm.DB) ILoanRepository {
	return &Loan{db}
}

func (l *Loan) Create(en *entity.Loan) error {
	return l.db.Create(en).Error
}
