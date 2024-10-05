package investment

import (
	repository "loan_engine/repository/investment"
)

type InvestmentServiceImpl struct {
	repository repository.Repository
}

func NewInvesmentService(r repository.Repository) InvestmentService {
	return &InvestmentServiceImpl{
		repository: r,
	}
}
