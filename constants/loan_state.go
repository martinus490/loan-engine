package constants

type LoanState string

const (
	Proposed  LoanState = "PROPOSED"
	Approved  LoanState = "APPROVED"
	Invested  LoanState = "INVESTED"
	Disbursed LoanState = "DISBURSED"
)
