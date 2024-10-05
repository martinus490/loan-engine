package test

import (
	"fmt"
	mocks "loan_engine/mocks/repository/investment"
	"loan_engine/model"
	"loan_engine/service/investment"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSaveInvestment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := investment.NewInvesmentService(mockRepo)

	currentTime := time.Now()
	investment := model.Investment{
		Id:               1,
		LoanId:           1,
		InvestorId:       1,
		Amount:           1000,
		CreatedAt:        currentTime,
		UpdatedAt:        currentTime,
		State:            "Open",
		ExpectedInterest: 50,
	}

	mockRepo.EXPECT().Save(investment).Return(investment, nil).Times(1)

	savedInvestment, err := service.Save(investment)

	assert.NoError(t, err, "Expected no error, but got one")
	assert.Equal(t, investment, savedInvestment, "Expected saved investment to match original")
}

func TestSaveInvestment_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := investment.NewInvesmentService(mockRepo)

	currentTime := time.Now()
	investment := model.Investment{
		Id:               1,
		LoanId:           1,
		InvestorId:       1,
		Amount:           1000,
		CreatedAt:        currentTime,
		UpdatedAt:        currentTime,
		State:            "Open",
		ExpectedInterest: 50,
	}

	mockRepo.EXPECT().Save(investment).Return(model.Investment{}, fmt.Errorf("save error")).Times(1)

	savedInvestment, err := service.Save(investment)

	assert.Error(t, err, "Expected an error when saving investment")
	assert.Equal(t, "save error", err.Error(), "Expected error message to match")
	assert.Equal(t, model.Investment{}, savedInvestment, "Expected saved investment to be empty")
}
