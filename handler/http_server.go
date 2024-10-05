package handler

import (
	loanService "loan_engine/service/loan"

	"github.com/go-playground/validator/v10"
)

type HttpServer struct {
	LoanService loanService.LoanService
	Validator   *validator.Validate
}

func NewHttpServer(
	loanService loanService.LoanService,
) *HttpServer {
	return &HttpServer{
		LoanService: loanService,
		Validator:   validator.New(),
	}
}
