package loan

import (
	"fmt"
	"loan_engine/constants"
	"loan_engine/model"
	types "loan_engine/types"
	"time"
)

func (s *LoanServiceImpl) GetPaginatedLoans(page int, pageSize int) (*types.PaginatedLoansResponse, error) {
	loans, totalRecords, err := s.repository.FetchLoansWithPagination(page, pageSize)
	if err != nil {
		return nil, err
	}

	response := types.PaginatedLoansResponse{
		Data:         loans,
		TotalRecords: totalRecords,
		Page:         page,
		PageSize:     pageSize,
	}

	return &response, err
}

func (s *LoanServiceImpl) GetLoanById(id int64) (model.Loan, error) {
	return s.repository.FindById(id)
}

func (s *LoanServiceImpl) CreateLoan(request types.CreateLoanRequest) (*types.CreateLoanResponse, error) {
	newLoan := model.Loan{
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		BorrowerId:      request.BorrowerId,
		PrincipalAmount: request.PrincipalAmount,
		Rate:            request.Rate,
		MaturityDate:    request.MaturityDate,
		State:           string(constants.Proposed),
	}

	createdLoan, err := s.repository.Save(newLoan)
	if err != nil {
		return nil, err
	}

	response := types.CreateLoanResponse{
		Id:              createdLoan.Id,
		CreatedAt:       createdLoan.CreatedAt,
		UpdatedAt:       createdLoan.UpdatedAt,
		BorrowerId:      createdLoan.BorrowerId,
		PrincipalAmount: createdLoan.PrincipalAmount,
		Rate:            createdLoan.Rate,
		MaturityDate:    createdLoan.MaturityDate,
		State:           createdLoan.State,
	}

	return &response, nil
}

func (s *LoanServiceImpl) ApprovalLoan(request types.ApproveLoanRequest) (*types.ApproveLoanResponse, error) {
	existingLoan, err := s.repository.GetById(request.LoanId)
	if err != nil {
		return nil, fmt.Errorf("loan id %d doesnt't exists", request.LoanId)
	}

	newLoanApproval := model.LoanApproval{
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
		LoanId:                   request.LoanId,
		ApprovalDate:             time.Now(),
		FieldValidatorEmployeeId: request.FieldValidatorEmployeeId,
		FieldValidatePicture:     request.FieldValidatePicture,
	}

	loanApproval, err := s.loanApprovalService.Save(newLoanApproval)
	if err != nil {
		return nil, err
	}

	existingLoan.State = string(constants.Approved)
	existingLoan, err = s.repository.Update(existingLoan)
	if err != nil {
		return nil, err
	}

	response := types.ApproveLoanResponse{
		Id:                       loanApproval.Id,
		ApprovalDate:             loanApproval.ApprovalDate,
		Loan:                     existingLoan,
		FieldValidatorEmployeeId: loanApproval.FieldValidatorEmployeeId,
		FieldValidatePicture:     loanApproval.FieldValidatePicture,
	}

	return &response, nil
}

func (s *LoanServiceImpl) LoanFunding(request types.FundingRequest) (*types.FundingResponse, error) {
	existingLoan, err := s.repository.GetById(request.LoanId)
	if err != nil {
		return nil, fmt.Errorf("loan id %d doesnt't exists", request.LoanId)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if existingLoan.PrincipalAmount < existingLoan.InvestedAmount+request.Amount {
		return nil, fmt.Errorf("funding amount must be smaller than current loan invested amount")
	}

	newInvestment := model.Investment{
		LoanId:           request.LoanId,
		InvestorId:       request.InvestorId,
		Amount:           request.Amount,
		ExpectedInterest: request.Amount * existingLoan.Rate / 100,
		State:            string(constants.Open),
	}

	investment, err := s.investmentService.Save(newInvestment)
	if err != nil {
		return nil, err
	}

	existingLoan.InvestedAmount += request.Amount
	if existingLoan.InvestedAmount == existingLoan.PrincipalAmount {
		existingLoan.State = string(constants.Invested)
	}

	existingLoan, err = s.repository.Update(existingLoan)
	if err != nil {
		return nil, err
	}

	response := types.FundingResponse{
		Id:         investment.Id,
		Loan:       existingLoan,
		InvestorId: investment.InvestorId,
		Amount:     investment.Amount,
		State:      investment.State,
	}

	return &response, nil
}

func (s *LoanServiceImpl) LoanDisburse(request types.DisbursementRequest) (*types.DisbursementResponse, error) {
	existingLoan, err := s.repository.GetById(request.LoanId)
	if err != nil {
		return nil, fmt.Errorf("loan id %d doesnt't exists", request.LoanId)
	}

	if existingLoan.State != string(constants.Invested) {
		return nil, fmt.Errorf("invested amount for loan id %d is still not enough", request.LoanId)
	}

	newDisbursement := model.Disbursement{
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
		LoanId:                 request.LoanId,
		DisbursementDate:       time.Now(),
		AgreementLetter:        request.AgreementLetter,
		FieldOfficerEmployeeId: request.FieldOfficerEmployeeId,
	}

	disbursement, err := s.disbursementService.Save(newDisbursement)
	if err != nil {
		return nil, err
	}

	existingLoan.State = string(constants.Disbursed)
	existingLoan, err = s.repository.Update(existingLoan)
	if err != nil {
		return nil, err
	}

	response := types.DisbursementResponse{
		Id:                     disbursement.Id,
		Loan:                   existingLoan,
		DisbursementDate:       disbursement.DisbursementDate,
		AgreementLetter:        disbursement.AgreementLetter,
		FieldOfficerEmployeeId: disbursement.FieldOfficerEmployeeId,
	}

	return &response, nil
}
