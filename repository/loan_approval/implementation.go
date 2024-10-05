package loan_approval

import "loan_engine/model"

func (r *loanApprovalRepository) Save(loanApproval model.LoanApproval) (model.LoanApproval, error) {
	if err := r.db.Create(&loanApproval).Error; err != nil {
		return model.LoanApproval{}, err
	}

	return loanApproval, nil
}
