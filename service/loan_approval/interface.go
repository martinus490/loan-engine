package loan_appoval

import (
	"loan_engine/model"
)

type LoanApprovalService interface {
	Save(loanApproval model.LoanApproval) (model.LoanApproval, error)
}
