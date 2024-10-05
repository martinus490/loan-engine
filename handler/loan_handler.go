package handler

import (
	"loan_engine/types"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *HttpServer) GetLoansWithPagination(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	response, err := h.LoanService.GetPaginatedLoans(page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, constructErrorReponse("Error get paginated loan", err.Error()))
	}

	return c.JSON(http.StatusOK, response)
}

func (s *HttpServer) GetLoanById(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, constructErrorReponse("Invalid loan id", err.Error()))
	}

	loan, err := s.LoanService.GetLoanById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(http.StatusNotFound, constructErrorReponse("Loan not found", err.Error()))
		}
		return ctx.JSON(http.StatusInternalServerError, constructErrorReponse("Internal server error", err.Error()))
	}

	return ctx.JSON(http.StatusOK, loan)
}

func (s *HttpServer) CreteLoan(ctx echo.Context) error {
	var request types.CreateLoanRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println("[CreteLoan] invalid request body", err)
		return ctx.JSON(http.StatusBadRequest, constructErrorReponse("Invalid request body", err.Error()))
	}

	if err := s.Validator.Struct(request); err != nil {
		log.Println("[CreteLoan] invalid request body", err)
		return ctx.JSON(http.StatusBadRequest, constructErrorReponse("Validation error", err.Error()))
	}

	response, err := s.LoanService.CreateLoan(request)
	if err != nil {
		log.Println("[CreteLoan] error create loan", err)
		return ctx.JSON(http.StatusInternalServerError, constructErrorReponse("Error creating loan", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response)
}

func (s *HttpServer) ApproveLoan(ctx echo.Context) error {
	var request types.ApproveLoanRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println("[ApproveLoan] invalid request body", err)
		return ctx.JSON(http.StatusBadRequest, constructErrorReponse("Invalid request body", err.Error()))
	}

	if err := s.Validator.Struct(request); err != nil {
		log.Println("[ApproveLoan] invalid request body", err)
		return ctx.JSON(http.StatusBadRequest, constructErrorReponse("Validation error", err.Error()))
	}

	response, err := s.LoanService.ApprovalLoan(request)
	if err != nil {
		log.Println("[ApproveLoan] error approve loan", err)
		return ctx.JSON(http.StatusInternalServerError, constructErrorReponse("Error approve loan", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response)
}

func (s *HttpServer) LoanFunding(ctx echo.Context) error {
	var request types.FundingRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println("[Funding] invalid request body", err)
		return ctx.JSON(http.StatusBadRequest, constructErrorReponse("Invalid request body", err.Error()))
	}

	if err := s.Validator.Struct(request); err != nil {
		log.Println("[Funding] invalid request body", err)
		return ctx.JSON(http.StatusBadRequest, constructErrorReponse("Validation error", err.Error()))
	}

	response, err := s.LoanService.LoanFunding(request)
	if err != nil {
		log.Println("[Funding] error funding loan", err)
		return ctx.JSON(http.StatusInternalServerError, constructErrorReponse("Error funding loan", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response)
}

func (s *HttpServer) LoanDisburse(ctx echo.Context) error {
	var request types.DisbursementRequest

	if err := ctx.Bind(&request); err != nil {
		log.Println("[Disburse] invalid request body", err)
		return ctx.JSON(http.StatusBadRequest, constructErrorReponse("Invalid request body", err.Error()))
	}

	if err := s.Validator.Struct(request); err != nil {
		log.Println("[Disburse] invalid request body", err)
		return ctx.JSON(http.StatusBadRequest, constructErrorReponse("Validation error", err.Error()))
	}

	response, err := s.LoanService.LoanDisburse(request)
	if err != nil {
		log.Println("[Disburse] error disburse loan", err)
		return ctx.JSON(http.StatusInternalServerError, constructErrorReponse("Error disburse loan", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response)
}

func constructErrorReponse(message string, detail string) types.ErrorResponse {
	return types.ErrorResponse{
		Message: message,
		Detail:  detail,
	}
}
