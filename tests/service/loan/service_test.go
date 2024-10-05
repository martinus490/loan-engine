package test

import (
	"fmt"
	"testing"
	"time"

	"loan_engine/constants"
	mocks "loan_engine/mocks/repository/loan"
	disbursementServiceMock "loan_engine/mocks/service/disbursement"
	investmentServiceMock "loan_engine/mocks/service/investment"
	loanApprovalServiceMock "loan_engine/mocks/service/loan_approval"
	"loan_engine/model"
	"loan_engine/service/loan"
	types "loan_engine/types"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateLoan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockLoanApprovalService := loanApprovalServiceMock.NewMockLoanApprovalService(ctrl)
	mockInvestmentService := investmentServiceMock.NewMockInvestmentService(ctrl)
	mockDisbursementService := disbursementServiceMock.NewMockDisbursementService(ctrl)
	service := loan.NewLoanService(mockRepo, mockLoanApprovalService, mockInvestmentService, mockDisbursementService)

	request := types.CreateLoanRequest{
		BorrowerId:      1,
		PrincipalAmount: 1000,
		Rate:            5,
		MaturityDate:    time.Now().AddDate(1, 0, 0),
	}

	expectedLoan := model.Loan{
		Id:              1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		BorrowerId:      request.BorrowerId,
		PrincipalAmount: request.PrincipalAmount,
		Rate:            request.Rate,
		MaturityDate:    request.MaturityDate,
		State:           string(constants.Proposed),
	}

	mockRepo.EXPECT().Save(gomock.Any()).Return(expectedLoan, nil).Times(1)

	response, err := service.CreateLoan(request)

	assert.NoError(t, err, "Expected no error, but got one")

	assert.Equal(t, expectedLoan.PrincipalAmount, response.PrincipalAmount, "Expected PrincipalAmount to match")
	assert.Equal(t, expectedLoan.State, response.State, "Expected State to match")
	assert.Equal(t, expectedLoan.BorrowerId, response.BorrowerId, "Expected BorrowerId to match")
	assert.Equal(t, expectedLoan.Rate, response.Rate, "Expected Rate to match")
	assert.Equal(t, expectedLoan.MaturityDate, response.MaturityDate, "Expected MaturityDate to match")
	assert.Equal(t, expectedLoan.CreatedAt, response.CreatedAt, "Expected CreatedAt to match")
	assert.Equal(t, expectedLoan.UpdatedAt, response.UpdatedAt, "Expected UpdatedAt to match")
	assert.Equal(t, expectedLoan.Id, response.Id, "Expected Id to match")
}

func TestGetLoanById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockLoanApprovalService := loanApprovalServiceMock.NewMockLoanApprovalService(ctrl)
	mockInvestmentService := investmentServiceMock.NewMockInvestmentService(ctrl)
	mockDisbursementService := disbursementServiceMock.NewMockDisbursementService(ctrl)

	service := loan.NewLoanService(mockRepo, mockLoanApprovalService, mockInvestmentService, mockDisbursementService)

	expectedLoan := model.Loan{Id: 1, BorrowerId: 1, PrincipalAmount: 1000}

	mockRepo.EXPECT().FindById(int64(1)).Return(expectedLoan, nil).Times(1)

	loan, err := service.GetLoanById(1)

	assert.NoError(t, err, "Expected no error, but got one")
	assert.Equal(t, expectedLoan.Id, loan.Id, "Expected loan ID to match")
	assert.Equal(t, expectedLoan.BorrowerId, loan.BorrowerId, "Expected loan BorrowerId to match")
	assert.Equal(t, expectedLoan.PrincipalAmount, loan.PrincipalAmount, "Expected loan PrincipalAmount to match")
}

func TestApprovalLoan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockLoanApprovalService := loanApprovalServiceMock.NewMockLoanApprovalService(ctrl)
	mockInvestmentService := investmentServiceMock.NewMockInvestmentService(ctrl)
	mockDisbursementService := disbursementServiceMock.NewMockDisbursementService(ctrl)

	loanService := loan.NewLoanService(mockRepo, mockLoanApprovalService, mockInvestmentService, mockDisbursementService)

	request := types.ApproveLoanRequest{
		LoanId:                   1,
		FieldValidatorEmployeeId: 1,
		FieldValidatePicture:     "picture.jpg",
	}

	mockRepo.EXPECT().
		GetById(request.LoanId).
		Return(model.Loan{
			Id:              request.LoanId,
			State:           string(constants.Proposed),
			PrincipalAmount: 1000,
			InvestedAmount:  0,
		}, nil)

	mockLoanApprovalService.EXPECT().
		Save(gomock.Any()).
		Return(model.LoanApproval{
			Id:                       1,
			ApprovalDate:             time.Now(),
			FieldValidatorEmployeeId: request.FieldValidatorEmployeeId,
			FieldValidatePicture:     request.FieldValidatePicture,
		}, nil)

	mockRepo.EXPECT().
		Update(gomock.Any()).
		Return(model.Loan{
			Id:              request.LoanId,
			State:           string(constants.Approved),
			PrincipalAmount: 1000,
			InvestedAmount:  0,
		}, nil)

	response, err := loanService.ApprovalLoan(request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, request.LoanId, response.Loan.Id)
	assert.Equal(t, string(constants.Approved), response.Loan.State)
}

func TestLoanFunding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockLoanApprovalService := loanApprovalServiceMock.NewMockLoanApprovalService(ctrl)
	mockInvestmentService := investmentServiceMock.NewMockInvestmentService(ctrl)
	mockDisbursementService := disbursementServiceMock.NewMockDisbursementService(ctrl)
	service := loan.NewLoanService(mockRepo, mockLoanApprovalService, mockInvestmentService, mockDisbursementService)

	existingLoan := model.Loan{
		Id:              1,
		BorrowerId:      1,
		PrincipalAmount: 1000,
		InvestedAmount:  200,
		Rate:            5,
		State:           string(constants.Open),
	}

	request := types.FundingRequest{
		LoanId:     existingLoan.Id,
		InvestorId: 2,
		Amount:     300,
	}

	mockRepo.EXPECT().GetById(request.LoanId).Return(existingLoan, nil).Times(1)

	expectedInvestment := model.Investment{
		Id:               1,
		LoanId:           request.LoanId,
		InvestorId:       request.InvestorId,
		Amount:           request.Amount,
		ExpectedInterest: request.Amount * existingLoan.Rate / 100,
		State:            string(constants.Open),
	}

	mockInvestmentService.EXPECT().Save(gomock.Any()).Return(expectedInvestment, nil).Times(1)

	existingLoan.InvestedAmount += request.Amount
	mockRepo.EXPECT().Update(existingLoan).Return(existingLoan, nil).Times(1)

	response, err := service.LoanFunding(request)

	assert.NoError(t, err, "Expected no error, but got one")

	assert.Equal(t, expectedInvestment.Id, response.Id, "Expected Investment Id to match")
	assert.Equal(t, existingLoan, response.Loan, "Expected Loan to match")
	assert.Equal(t, expectedInvestment.InvestorId, response.InvestorId, "Expected InvestorId to match")
	assert.Equal(t, expectedInvestment.Amount, response.Amount, "Expected Amount to match")
	assert.Equal(t, expectedInvestment.State, response.State, "Expected State to match")
}

func TestLoanFunding_AmountExceedsInvestedAmount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockLoanApprovalService := loanApprovalServiceMock.NewMockLoanApprovalService(ctrl)
	mockInvestmentService := investmentServiceMock.NewMockInvestmentService(ctrl)
	mockDisbursementService := disbursementServiceMock.NewMockDisbursementService(ctrl)
	service := loan.NewLoanService(mockRepo, mockLoanApprovalService, mockInvestmentService, mockDisbursementService)

	existingLoan := model.Loan{
		Id:              1,
		BorrowerId:      1,
		PrincipalAmount: 1000,
		InvestedAmount:  800,
		Rate:            5,
		State:           string(constants.Open),
	}

	request := types.FundingRequest{
		LoanId:     existingLoan.Id,
		InvestorId: 2,
		Amount:     300,
	}

	mockRepo.EXPECT().GetById(request.LoanId).Return(existingLoan, nil).Times(1)

	response, err := service.LoanFunding(request)

	assert.Error(t, err, "Expected error for funding amount exceeding limit")

	assert.Nil(t, response, "Expected response to be nil")
	assert.Equal(t, "funding amount must be smaller than current loan invested amount", err.Error())
}

func TestLoanDisburse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockDisbursementService := disbursementServiceMock.NewMockDisbursementService(ctrl)
	service := loan.NewLoanService(mockRepo, nil, nil, mockDisbursementService)

	existingLoan := model.Loan{
		Id:              1,
		BorrowerId:      1,
		PrincipalAmount: 1000,
		InvestedAmount:  1000,
		Rate:            5,
		State:           string(constants.Invested),
	}

	request := types.DisbursementRequest{
		LoanId:                 existingLoan.Id,
		AgreementLetter:        "agreement.pdf",
		FieldOfficerEmployeeId: 1,
	}

	mockRepo.EXPECT().GetById(request.LoanId).Return(existingLoan, nil).Times(1)

	expectedDisbursement := model.Disbursement{
		Id:                     1,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
		LoanId:                 request.LoanId,
		DisbursementDate:       time.Now(),
		AgreementLetter:        request.AgreementLetter,
		FieldOfficerEmployeeId: request.FieldOfficerEmployeeId,
	}

	mockDisbursementService.EXPECT().Save(gomock.Any()).Return(expectedDisbursement, nil).Times(1)

	existingLoan.State = string(constants.Disbursed)
	mockRepo.EXPECT().Update(existingLoan).Return(existingLoan, nil).Times(1)

	response, err := service.LoanDisburse(request)

	assert.NoError(t, err, "Expected no error, but got one")
	assert.Equal(t, expectedDisbursement.Id, response.Id, "Expected Disbursement Id to match")
	assert.Equal(t, existingLoan, response.Loan, "Expected Loan to match")
	assert.Equal(t, expectedDisbursement.DisbursementDate, response.DisbursementDate, "Expected DisbursementDate to match")
	assert.Equal(t, expectedDisbursement.AgreementLetter, response.AgreementLetter, "Expected AgreementLetter to match")
	assert.Equal(t, expectedDisbursement.FieldOfficerEmployeeId, response.FieldOfficerEmployeeId, "Expected FieldOfficerEmployeeId to match")
}

func TestLoanDisburse_LoanNotInvested(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockDisbursementService := disbursementServiceMock.NewMockDisbursementService(ctrl)
	service := loan.NewLoanService(mockRepo, nil, nil, mockDisbursementService)

	existingLoan := model.Loan{
		Id:              1,
		BorrowerId:      1,
		PrincipalAmount: 1000,
		InvestedAmount:  800,
		Rate:            5,
		State:           string(constants.Open),
	}

	request := types.DisbursementRequest{
		LoanId:                 existingLoan.Id,
		AgreementLetter:        "agreement.pdf",
		FieldOfficerEmployeeId: 1,
	}

	mockRepo.EXPECT().GetById(request.LoanId).Return(existingLoan, nil).Times(1)

	response, err := service.LoanDisburse(request)

	assert.Error(t, err, "Expected an error when loan is not in 'Invested' state")
	assert.Nil(t, response, "Expected response to be nil")
	assert.Equal(t, "invested amount for loan id 1 is still not enough", err.Error(), "Expected specific error message")
}

func TestLoanDisburse_LoanNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockDisbursementService := disbursementServiceMock.NewMockDisbursementService(ctrl)
	service := loan.NewLoanService(mockRepo, nil, nil, mockDisbursementService)

	request := types.DisbursementRequest{
		LoanId:                 999,
		AgreementLetter:        "agreement.pdf",
		FieldOfficerEmployeeId: 1,
	}

	mockRepo.EXPECT().GetById(request.LoanId).Return(model.Loan{}, fmt.Errorf("loan not found")).Times(1)

	response, err := service.LoanDisburse(request)

	assert.Error(t, err, "Expected error for non-existent loan ID")
	assert.Nil(t, response, "Expected response to be nil")
	assert.Equal(t, "loan id 999 doesnt't exists", err.Error())
}

func TestLoanDisburse_SaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockDisbursementService := disbursementServiceMock.NewMockDisbursementService(ctrl)
	service := loan.NewLoanService(mockRepo, nil, nil, mockDisbursementService)

	existingLoan := model.Loan{
		Id:              1,
		BorrowerId:      1,
		PrincipalAmount: 1000,
		InvestedAmount:  1000,
		Rate:            5,
		State:           string(constants.Invested),
	}

	request := types.DisbursementRequest{
		LoanId:                 existingLoan.Id,
		AgreementLetter:        "agreement.pdf",
		FieldOfficerEmployeeId: 1,
	}

	mockRepo.EXPECT().GetById(request.LoanId).Return(existingLoan, nil).Times(1)

	mockDisbursementService.EXPECT().Save(gomock.Any()).Return(model.Disbursement{}, fmt.Errorf("failed to save disbursement")).Times(1)

	response, err := service.LoanDisburse(request)

	assert.Error(t, err, "Expected error for failed disbursement save")
	assert.Nil(t, response, "Expected response to be nil")
	assert.Equal(t, "failed to save disbursement", err.Error())
}

func TestLoanDisburse_UpdateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	mockDisbursementService := disbursementServiceMock.NewMockDisbursementService(ctrl)
	service := loan.NewLoanService(mockRepo, nil, nil, mockDisbursementService)

	existingLoan := model.Loan{
		Id:              1,
		BorrowerId:      1,
		PrincipalAmount: 1000,
		InvestedAmount:  1000,
		Rate:            5,
		State:           string(constants.Invested),
	}

	request := types.DisbursementRequest{
		LoanId:                 existingLoan.Id,
		AgreementLetter:        "agreement.pdf",
		FieldOfficerEmployeeId: 1,
	}

	mockRepo.EXPECT().GetById(request.LoanId).Return(existingLoan, nil).Times(1)

	mockDisbursementService.EXPECT().Save(gomock.Any()).Return(model.Disbursement{}, nil).Times(1)

	existingLoan.State = string(constants.Disbursed)
	mockRepo.EXPECT().Update(existingLoan).Return(model.Loan{}, fmt.Errorf("failed to update loan")).Times(1)

	response, err := service.LoanDisburse(request)

	assert.Error(t, err, "Expected error for failed loan update")
	assert.Nil(t, response, "Expected response to be nil")
	assert.Equal(t, "failed to update loan", err.Error())
}
