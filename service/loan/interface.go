package loan

import (
	"loan_engine/model"
	types "loan_engine/types"
)

type LoanService interface {
	GetPaginatedLoans(page int, pageSize int) (*types.PaginatedLoansResponse, error)
	GetLoanById(id int64) (model.Loan, error)
	CreateLoan(request types.CreateLoanRequest) (*types.CreateLoanResponse, error)
	ApprovalLoan(request types.ApproveLoanRequest) (*types.ApproveLoanResponse, error)
	LoanFunding(request types.FundingRequest) (*types.FundingResponse, error)
	LoanDisburse(request types.DisbursementRequest) (*types.DisbursementResponse, error)
}
