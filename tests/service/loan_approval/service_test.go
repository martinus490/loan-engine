package test

import (
	"fmt"
	mocks "loan_engine/mocks/repository/loan_approval"
	"loan_engine/model"
	loan_appoval "loan_engine/service/loan_approval"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSaveLoanApproval(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := loan_appoval.NewLoanApprovalService(mockRepo)

	loanApproval := model.LoanApproval{
		Id:                       1,
		LoanId:                   1,
		ApprovalDate:             time.Now(),
		FieldValidatorEmployeeId: 123,
		FieldValidatePicture:     "picture.jpg",
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}

	mockRepo.EXPECT().Save(loanApproval).Return(loanApproval, nil).Times(1)

	savedLoanApproval, err := service.Save(loanApproval)

	assert.NoError(t, err, "Expected no error, but got one")

	assert.Equal(t, loanApproval, savedLoanApproval, "Expected saved loan approval to match original")
}

func TestSaveLoanApproval_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := loan_appoval.NewLoanApprovalService(mockRepo)

	loanApproval := model.LoanApproval{
		Id:                       1,
		LoanId:                   1,
		ApprovalDate:             time.Now(),
		FieldValidatorEmployeeId: 123,
		FieldValidatePicture:     "picture.jpg",
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}

	mockRepo.EXPECT().Save(loanApproval).Return(model.LoanApproval{}, fmt.Errorf("save error")).Times(1)

	savedLoanApproval, err := service.Save(loanApproval)

	assert.Error(t, err, "Expected an error when saving loan approval")
	assert.Equal(t, "save error", err.Error(), "Expected error message to match")

	assert.Equal(t, model.LoanApproval{}, savedLoanApproval, "Expected saved loan approval to be empty")
}
