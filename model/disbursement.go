package model

import "time"

type Disbursement struct {
	Id                     int64     `gorm:"primaryKey"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
	UpdatedAt              time.Time `gorm:"autoUpdateTime"`
	LoanId                 int64     `gorm:"not null"`
	DisbursementDate       time.Time `gorm:"not null"`
	AgreementLetter        string    `gorm:"not null"`
	FieldOfficerEmployeeId int64     `gorm:"not null"`
}

func (Disbursement) TableName() string {
	return "disbursement"
}
