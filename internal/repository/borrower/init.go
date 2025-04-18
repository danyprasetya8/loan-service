//go:generate mockgen -destination=../../../mocks/mock_borrower_repository.go -package=mocks loan-service/internal/repository/borrower IBorrowerRepository
package borrower

import (
	"loan-service/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IBorrowerRepository interface {
	Create(en *entity.Borrower) error
	Count() int64
	GetByID(id string) *entity.Borrower
	GetList(page, size int) []entity.Borrower
	Delete(id string) (bool, error)
}

type Borrower struct {
	db *gorm.DB
}

func New(db *gorm.DB) IBorrowerRepository {
	return &Borrower{db}
}

func (b *Borrower) Create(en *entity.Borrower) error {
	en.ID = uuid.New().String()
	return b.db.Create(en).Error
}

func (b *Borrower) Count() (c int64) {
	b.db.Model(&entity.Borrower{}).
		Count(&c)
	return
}

func (b *Borrower) GetByID(id string) *entity.Borrower {
	borrower := &entity.Borrower{}

	err := b.db.Where("id = ?", id).
		First(borrower).
		Error

	if err != nil {
		return nil
	}

	return borrower
}

func (b *Borrower) GetList(page, size int) (borrowers []entity.Borrower) {
	borrowers = make([]entity.Borrower, 0)

	b.db.Limit(size).
		Offset((page - 1) * size).
		Find(&borrowers)
	return
}

func (b *Borrower) Delete(id string) (bool, error) {
	result := b.db.Delete(&entity.Borrower{}, "id = ?", id)

	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}
