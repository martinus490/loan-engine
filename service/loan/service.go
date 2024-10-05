package loan

import (
	repository "loan_engine/repository/loan"
	disbursementService "loan_engine/service/disbursement"
	investmentService "loan_engine/service/investment"
	loanApprovalService "loan_engine/service/loan_approval"
	"sync"
)

type LoanServiceImpl struct {
	repository          repository.Repository
	loanApprovalService loanApprovalService.LoanApprovalService
	investmentService   investmentService.InvestmentService
	disbursementService disbursementService.DisbursementService
	mu                  sync.Mutex
}

func NewLoanService(
	r repository.Repository,
	loanApprovalService loanApprovalService.LoanApprovalService,
	investmentService investmentService.InvestmentService,
	disbursementService disbursementService.DisbursementService,
) LoanService {
	return &LoanServiceImpl{
		repository:          r,
		loanApprovalService: loanApprovalService,
		investmentService:   investmentService,
		disbursementService: disbursementService,
	}
}
