package investment

import "loan_engine/model"

func (r *investmentRepository) Save(investment model.Investment) (model.Investment, error) {
	if err := r.db.Create(&investment).Error; err != nil {
		return model.Investment{}, err
	}

	return investment, nil
}
