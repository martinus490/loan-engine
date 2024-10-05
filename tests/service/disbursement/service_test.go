package test

import (
	"fmt"
	mocks "loan_engine/mocks/repository/disbursement"
	"loan_engine/model"
	"loan_engine/service/disbursement"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSaveDisbursement_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := disbursement.NewDisbursementService(mockRepo)

	currentTime := time.Now()
	disbursement := model.Disbursement{
		Id:                     1,
		LoanId:                 1,
		DisbursementDate:       currentTime,
		AgreementLetter:        "agreement.pdf",
		FieldOfficerEmployeeId: 123,
		CreatedAt:              currentTime,
		UpdatedAt:              currentTime,
	}

	mockRepo.EXPECT().Save(disbursement).Return(disbursement, nil).Times(1)

	savedDisbursement, err := service.Save(disbursement)

	assert.NoError(t, err, "Expected no error, but got one")
	assert.Equal(t, disbursement, savedDisbursement, "Expected saved disbursement to match original")
}

func TestSaveDisbursement_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := disbursement.NewDisbursementService(mockRepo)

	currentTime := time.Now()
	disbursement := model.Disbursement{
		Id:                     1,
		LoanId:                 1,
		DisbursementDate:       currentTime,
		AgreementLetter:        "agreement.pdf",
		FieldOfficerEmployeeId: 123,
		CreatedAt:              currentTime,
		UpdatedAt:              currentTime,
	}

	mockRepo.EXPECT().Save(disbursement).Return(model.Disbursement{}, fmt.Errorf("save error")).Times(1)

	savedDisbursement, err := service.Save(disbursement)

	assert.Error(t, err, "Expected an error when saving disbursement")
	assert.Equal(t, "save error", err.Error(), "Expected error message to match")
	assert.Equal(t, model.Disbursement{}, savedDisbursement, "Expected saved disbursement to be empty")
}
