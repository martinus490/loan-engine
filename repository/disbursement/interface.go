package disbursement

import "loan_engine/model"

type Repository interface {
	Save(model.Disbursement) (model.Disbursement, error)
}
