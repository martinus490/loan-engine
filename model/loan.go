package model

import "time"

type Loan struct {
	Id              int64     `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	BorrowerId      int64     `json:"borrower_id" gorm:"not null"`
	PrincipalAmount float64   `json:"principal_amount" gorm:"not null"`
	Rate            float64   `json:"rate" gorm:"not null"`
	State           string    `json:"state" gorm:"default:'PROPOSED'"`
	MaturityDate    time.Time `json:"maturity_date" gorm:"not null"`
	InvestedAmount  float64   `json:"invested_amount" form:"default:0"`
}

func (Loan) TableName() string {
	return "loan"
}
