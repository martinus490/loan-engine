package investment

import "loan_engine/model"

func (s *InvestmentServiceImpl) Save(investment model.Investment) (model.Investment, error) {
	return s.repository.Save(investment)
}
