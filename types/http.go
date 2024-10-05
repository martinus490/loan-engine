package types

import (
	"loan_engine/model"
	"time"
)

type PaginatedLoansResponse struct {
	Data         []model.Loan `json:"data"`
	TotalRecords int64        `json:"total_records"`
	Page         int          `json:"page"`
	PageSize     int          `json:"page_size"`
}

type CreateLoanRequest struct {
	BorrowerId      int64     `json:"borrower_id" validate:"required"`
	PrincipalAmount float64   `json:"principal_amount" validate:"required,gt=0"`
	Rate            float64   `json:"rate" validate:"required,gt=0"`
	MaturityDate    time.Time `json:"maturity_date" validate:"required"`
}

type CreateLoanResponse struct {
	Id              int64     `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	BorrowerId      int64     `json:"borrower_id"`
	PrincipalAmount float64   `json:"principal_amount"`
	Rate            float64   `json:"rate"`
	MaturityDate    time.Time `json:"maturity_date"`
	State           string    `json:"state"`
}

type ApproveLoanRequest struct {
	LoanId                   int64  `json:"loan_id" validate:"required"`
	FieldValidatorEmployeeId int64  `json:"field_validator_employee_id" validate:"required"`
	FieldValidatePicture     string `json:"field_validate_picture" validate:"required"`
}

type ApproveLoanResponse struct {
	Id                       int64      `json:"id"`
	Loan                     model.Loan `json:"loan"`
	ApprovalDate             time.Time  `json:"approval_date"`
	FieldValidatorEmployeeId int64      `json:"field_validator_employee_id"`
	FieldValidatePicture     string     `json:"field_validate_picture"`
}

type FundingRequest struct {
	LoanId     int64   `json:"loan_id" validate:"required"`
	InvestorId int64   `json:"investor_id" validate:"required"`
	Amount     float64 `json:"amount" validate:"required,gt=0"`
}

type FundingResponse struct {
	Id         int64      `json:"id"`
	Loan       model.Loan `json:"loan"`
	InvestorId int64      `json:"investor_id"`
	Amount     float64    `json:"amount"`
	State      string     `json:"state"`
}

type DisbursementRequest struct {
	LoanId                 int64  `json:"loan_id" validate:"required"`
	AgreementLetter        string `json:"aggrement_letter" validate:"required"`
	FieldOfficerEmployeeId int64  `json:"field_officer_employee_id" validate:"required"`
}

type DisbursementResponse struct {
	Id                     int64      `json:"id"`
	Loan                   model.Loan `json:"loan"`
	DisbursementDate       time.Time  `json:"disbursement_date"`
	AgreementLetter        string     `json:"aggrement_letter"`
	FieldOfficerEmployeeId int64      `json:"field_officer_employee_id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}
