package loan_approval

import "gorm.io/gorm"

type loanApprovalRepository struct {
	db *gorm.DB
}

func NewLoanApprovalRepository(db *gorm.DB) Repository {
	return &loanApprovalRepository{
		db: db,
	}
}
