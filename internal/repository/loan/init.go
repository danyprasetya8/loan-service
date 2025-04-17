package loan

import (
	"loan-service/internal/constant"
	"loan-service/internal/entity"

	"gorm.io/gorm"
)

type ILoanRepository interface {
	GetList(status constant.LoanStatus, page, size int) []entity.Loan
	Count(status constant.LoanStatus) int64
	GetDetail(id string) *entity.Loan
	Get(id string) *entity.Loan
	Create(en *entity.Loan) error
	Save(en *entity.Loan) error
}

type Loan struct {
	db *gorm.DB
}

func New(db *gorm.DB) ILoanRepository {
	return &Loan{db}
}

func (b *Loan) Count(status constant.LoanStatus) (c int64) {
	tx := b.db.Model(&entity.Loan{})

	if status != "" {
		tx = tx.Where("status = ?", status)
	}

	tx.Count(&c)
	return
}

func (b *Loan) GetList(status constant.LoanStatus, page, size int) []entity.Loan {
	loans := make([]entity.Loan, 0)

	tx := b.db

	if status != "" {
		tx = tx.Where("status = ?", status)
	}

	tx.Limit(size).
		Offset((page - 1) * size).
		Find(&loans)

	return loans
}

func (l *Loan) GetDetail(id string) *entity.Loan {
	loan := &entity.Loan{}

	err := l.db.Preload("LoanApproval").
		Preload("LoanDisbursement").
		Preload("LoanInvestment").
		First(&loan, "id = ?", id).
		Error

	if err != nil {
		return nil
	}

	return loan
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

func (l *Loan) Save(en *entity.Loan) error {
	return l.db.Save(en).Error
}
