package controller

import (
	"fmt"
	"goGuestThePlace/helpers"
	"goGuestThePlace/models"
	"goGuestThePlace/services"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controllerTransaction struct {
	TransactionRepository services.TransactionService
}

func TransactionController(transactionService services.TransactionService) *controllerTransaction {
	return &controllerTransaction{transactionService}
}

func (cr *controllerTransaction) GetTransactions(c echo.Context) error {
	transactions, err := cr.TransactionRepository.GetTransaction()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Transactions Success", transactions))

}

func (cr *controllerTransaction) CreateTransaction(c echo.Context) error {
	request := new(models.CreateTransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	transaction := models.Transaction{
		Price:      request.Price,
		Status:     request.Status,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	response, err := cr.TransactionRepository.CreateTransaction(transaction)
	fmt.Println(response)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Create Transaction Success", response))
}
