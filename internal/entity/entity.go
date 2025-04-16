package entity

import (
	"time"

	"gorm.io/gorm"
)

type Audit struct {
	CreatedAt time.Time
	CreatedBy string `gorm:"type:varchar(255)"`
	UpdatedAt time.Time
	UpdatedBy string `gorm:"type:varchar(255)"`
}

type Borrower struct {
	ID        string         `gorm:"primaryKey;type:varchar(255)"`
	Name      string         `gorm:"varchar(255)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Audit
}
