package main

import (
	database "loan_engine/config"
	"loan_engine/handler"
	disbursementRepository "loan_engine/repository/disbursement"
	investmentRepository "loan_engine/repository/investment"
	loanRepository "loan_engine/repository/loan"
	loanApprovalRepository "loan_engine/repository/loan_approval"
	disbursementService "loan_engine/service/disbursement"
	investmentService "loan_engine/service/investment"
	loanService "loan_engine/service/loan"
	loanApprovalService "loan_engine/service/loan_approval"

	"github.com/labstack/echo/v4"
)

func main() {
	// init database
	database.ConnectDatabase()

	// inject the connection to repository
	disbursementRepository := disbursementRepository.NewDisbursementRepository(database.DB)
	investmentRepository := investmentRepository.NewInvestmentRepository(database.DB)
	loanApprovalRepository := loanApprovalRepository.NewLoanApprovalRepository(database.DB)
	loanRepository := loanRepository.NewLoanRepository(database.DB)

	// inject repository to service
	disbursementService := disbursementService.NewDisbursementService(disbursementRepository)
	investmentService := investmentService.NewInvesmentService(investmentRepository)
	loanApprovalService := loanApprovalService.NewLoanApprovalService(loanApprovalRepository)
	loanService := loanService.NewLoanService(loanRepository, loanApprovalService, investmentService, disbursementService)

	// inject service to handler
	httpServer := handler.NewHttpServer(loanService)

	e := echo.New()

	handler.RegisterLoanRoutes(e, httpServer)

	e.Logger.Fatal(e.Start(":8080"))
}
