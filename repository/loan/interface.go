package loan

import "loan_engine/model"

type Repository interface {
	FetchLoansWithPagination(page int, pageSize int) ([]model.Loan, int64, error)
	FindById(id int64) (loan model.Loan, err error)
	GetById(id int64) (model.Loan, error)
	Save(model.Loan) (model.Loan, error)
	Update(loan model.Loan) (model.Loan, error)
}
