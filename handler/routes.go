package handler

import (
	"github.com/labstack/echo/v4"
)

func RegisterLoanRoutes(e *echo.Echo, s *HttpServer) {
	e.GET("/loan/available", s.GetLoansWithPagination)
	e.GET("/loan/:id", s.GetLoanById)
	e.POST("/loan", s.CreteLoan)
	e.POST("/loan/approve", s.ApproveLoan)
	e.POST("/loan/funding", s.LoanFunding)
	e.POST("/loan/disburse", s.LoanDisburse)
}
