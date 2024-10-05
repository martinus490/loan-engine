package investment

import (
	"loan_engine/model"
)

type InvestmentService interface {
	Save(request model.Investment) (model.Investment, error)
}
