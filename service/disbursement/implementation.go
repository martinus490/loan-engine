package disbursement

import "loan_engine/model"

func (s *DisbursementServiceImpl) Save(disbursement model.Disbursement) (model.Disbursement, error) {
	return s.repository.Save(disbursement)
}
