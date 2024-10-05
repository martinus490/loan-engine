package model

import "time"

type LoanApproval struct {
	Id                       int64     `gorm:"primaryKey"`
	CreatedAt                time.Time `gorm:"autoCreateTime"`
	UpdatedAt                time.Time `gorm:"autoUpdateTime"`
	LoanId                   int64     `gorm:"not null"`
	ApprovalDate             time.Time `gorm:"not null"`
	FieldValidatorEmployeeId int64     `gorm:"not null"`
	FieldValidatePicture     string    `gorm:"not null"`
}

func (LoanApproval) TableName() string {
	return "loan_approval"
}
