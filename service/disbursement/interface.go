package disbursement

import (
	"loan_engine/model"
)

type DisbursementService interface {
	Save(disbursement model.Disbursement) (model.Disbursement, error)
}
