package disbursement

import "gorm.io/gorm"

type disbursementRepository struct {
	db *gorm.DB
}

func NewDisbursementRepository(db *gorm.DB) Repository {
	return &disbursementRepository{
		db: db,
	}
}
