package disbursement

import (
	repository "loan_engine/repository/disbursement"
)

type DisbursementServiceImpl struct {
	repository repository.Repository
}

func NewDisbursementService(r repository.Repository) DisbursementService {
	return &DisbursementServiceImpl{
		repository: r,
	}
}
