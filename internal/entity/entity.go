package entity

import (
	"loan-service/internal/constant"
	"time"

	"gorm.io/gorm"
)

type Audit struct {
	CreatedAt time.Time
	CreatedBy string `gorm:"type:varchar(255)"`
	UpdatedAt time.Time
	UpdatedBy string `gorm:"type:varchar(255)"`
}

type File struct {
	ID           string            `gorm:"primaryKey;type:varchar(255)"`
	OriginalName string            `gorm:"not null"`
	Path         string            `gorm:"not null"`
	MimeType     string            `gorm:"not null"`
	Type         constant.FileType `gorm:"not null"`
	DeletedAt    gorm.DeletedAt    `gorm:"index"`
	Audit
}

type User struct {
	Email     string            `gorm:"primaryKey;type:varchar(255)"`
	Role      constant.UserRole `gorm:"not null"`
	DeletedAt gorm.DeletedAt    `gorm:"index"`
	Audit
}

type Borrower struct {
	ID        string         `gorm:"primaryKey;type:varchar(255)"`
	Name      string         `gorm:"varchar(255)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Audit
}

type Loan struct {
	ID              string `gorm:"primaryKey;type:varchar(255)"`
	BorrowerID      string
	Borrower        Borrower            `gorm:"constraint:OnDelete:CASCADE"`
	Status          constant.LoanStatus `gorm:"not null"`
	PrincipalAmount int64               `gorm:"not null"`
	InvestedAmount  int64               `gorm:"not null"`
	Rate            float64             `gorm:"not null"`
	ROI             float64             `gorm:"not null"`
	DeletedAt       gorm.DeletedAt      `gorm:"index"`
	Audit
}

type LoanApproval struct {
	ID             string `gorm:"primaryKey;type:varchar(255)"`
	LoanID         string
	Loan           Loan           `gorm:"constraint:OnDelete:CASCADE"`
	FieldOfficerID string         `gorm:"varchar(255);not null"`
	ProofOfPicture string         `gorm:"varchar(255);not null"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Audit
}

type LoanInvestment struct {
	ID         string `gorm:"primaryKey;type:varchar(255)"`
	LoanID     string
	Loan       Loan           `gorm:"constraint:OnDelete:CASCADE"`
	InvestorID string         `gorm:"varchar(255);not null"`
	Amount     int64          `gorm:"not null"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Audit
}

type LoanDisbursement struct {
	ID                      string `gorm:"primaryKey;type:varchar(255)"`
	LoanID                  string
	Loan                    Loan           `gorm:"constraint:OnDelete:CASCADE"`
	FieldOfficerID          string         `gorm:"varchar(255);not null"`
	BorrowerAgreementLetter string         `gorm:"varchar(255);not null"`
	DeletedAt               gorm.DeletedAt `gorm:"index"`
	Audit
}
