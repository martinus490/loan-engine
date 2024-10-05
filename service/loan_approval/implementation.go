package loan_appoval

import "loan_engine/model"

func (s *LoanApprovalServiceImpl) Save(loanApproval model.LoanApproval) (model.LoanApproval, error) {
	return s.repository.Save(loanApproval)
}
