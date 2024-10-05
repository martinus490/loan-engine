package loan

import "gorm.io/gorm"

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) Repository {
	return &loanRepository{
		db: db,
	}
}
