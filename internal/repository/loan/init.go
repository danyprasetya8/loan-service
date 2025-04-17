package loan

import (
	"loan-service/internal/entity"

	"gorm.io/gorm"
)

type ILoanRepository interface {
	Get(id string) *entity.Loan
	Create(en *entity.Loan) error
}

type Loan struct {
	db *gorm.DB
}

func New(db *gorm.DB) ILoanRepository {
	return &Loan{db}
}

func (l *Loan) Get(id string) *entity.Loan {
	loan := &entity.Loan{}

	tx := l.db.Where("id = ?", id).
		First(&loan)

	if tx.Error != nil {
		return nil
	}

	return loan
}

func (l *Loan) Create(en *entity.Loan) error {
	return l.db.Create(en).Error
}
