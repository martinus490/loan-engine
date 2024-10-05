package disbursement

import "loan_engine/model"

func (r *disbursementRepository) Save(disbursement model.Disbursement) (model.Disbursement, error) {
	if err := r.db.Create(&disbursement).Error; err != nil {
		return model.Disbursement{}, err
	}

	return disbursement, nil
}
