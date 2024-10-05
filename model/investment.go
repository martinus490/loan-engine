package model

import "time"

type Investment struct {
	Id               int64     `gorm:"primaryKey"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
	LoanId           int64     `gorm:"not null"`
	InvestorId       int64     `gorm:"not null"`
	Amount           float64   `gorm:"not null"`
	ExpectedInterest float64   `gorm:"not null"`
	State            string    `json:"state" gorm:"default:'OPEN'"`
}

func (Investment) TableName() string {
	return "investment"
}
