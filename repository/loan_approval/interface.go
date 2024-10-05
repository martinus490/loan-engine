package loan_approval

import "loan_engine/model"

type Repository interface {
	Save(model.LoanApproval) (model.LoanApproval, error)
}
