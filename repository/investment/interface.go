package investment

import "loan_engine/model"

type Repository interface {
	Save(model.Investment) (model.Investment, error)
}
