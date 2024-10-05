package investment

import "gorm.io/gorm"

type investmentRepository struct {
	db *gorm.DB
}

func NewInvestmentRepository(db *gorm.DB) Repository {
	return &investmentRepository{
		db: db,
	}
}
