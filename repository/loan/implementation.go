package loan

import (
	"loan_engine/constants"
	"loan_engine/model"
)

func (r *loanRepository) FetchLoansWithPagination(page int, pageSize int) ([]model.Loan, int64, error) {
	var loans []model.Loan
	var totalRecords int64

	if err := r.db.Model(&model.Loan{}).Where("state = ?", string(constants.Approved)).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := r.db.Where("state = ?", string(constants.Approved)).Limit(pageSize).Offset(offset).Find(&loans).Error; err != nil {
		return nil, 0, err
	}

	return loans, totalRecords, nil
}

func (r *loanRepository) FindById(id int64) (loan model.Loan, err error) {
	if err = r.db.First(&loan, id).Error; err != nil {
		return loan, err
	}
	return loan, nil
}

func (r *loanRepository) GetById(id int64) (loan model.Loan, err error) {
	if err := r.db.First(&loan, id).Error; err != nil {
		return loan, err
	}

	return loan, nil
}

func (r *loanRepository) Save(loan model.Loan) (model.Loan, error) {
	if err := r.db.Create(&loan).Error; err != nil {
		return model.Loan{}, err
	}

	return loan, nil
}

func (r *loanRepository) Update(loan model.Loan) (model.Loan, error) {
	if err := r.db.Save(&loan).Error; err != nil {
		return model.Loan{}, err
	}

	return loan, nil
}
