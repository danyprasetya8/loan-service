package file

import (
	"loan-service/internal/entity"

	"gorm.io/gorm"
)

type IFileRepository interface {
	Get(id string) *entity.File
	Create(en *entity.File) error
}

type File struct {
	db *gorm.DB
}

func New(db *gorm.DB) IFileRepository {
	return &File{db}
}

func (f *File) Get(id string) *entity.File {
	file := &entity.File{}
	tx := f.db.Where("id = ?", id).
		First(file)

	if tx.Error != nil {
		return nil
	}

	return file
}

func (f *File) Create(en *entity.File) error {
	return f.db.Create(en).Error
}
