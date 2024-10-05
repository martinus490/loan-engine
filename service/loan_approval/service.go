package loan_appoval

import (
	repository "loan_engine/repository/loan_approval"
)

type LoanApprovalServiceImpl struct {
	repository repository.Repository
}

func NewLoanApprovalService(r repository.Repository) LoanApprovalService {
	return &LoanApprovalServiceImpl{
		repository: r,
	}
}
