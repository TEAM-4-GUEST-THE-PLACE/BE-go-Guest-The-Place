package controller

import (
	"fmt"
	"goGuestThePlace/helpers"
	"goGuestThePlace/models"
	"goGuestThePlace/services"
	"net/http"
	"strconv"
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

func (cr *controllerTransaction) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	transactionStatus, ok := notificationPayload["transaction_status"].(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("invalid transaction_status"))
	}

	orderID, ok := notificationPayload["order_id"].(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("invalid order_id"))
	}

	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("invalid order_id format"))
	}

	fmt.Println("test payload", notificationPayload)

	switch transactionStatus {
	case "capture":
		fraudStatus, ok := notificationPayload["fraud_status"].(string)
		if !ok {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse("invalid fraud_status"))
		}
		switch fraudStatus {
		case "accept":
			cr.TransactionRepository.UpadateTransaction("success", orderIDInt)
		case "challenge":
			cr.TransactionRepository.UpadateTransaction("pending", orderIDInt)
		}
	case "settlement":
		cr.TransactionRepository.UpadateTransaction("success", orderIDInt)
	case "deny", "cancel", "expire":
		cr.TransactionRepository.UpadateTransaction("failed", orderIDInt)
	case "pending":
		cr.TransactionRepository.UpadateTransaction("pending", orderIDInt)
	default:
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("invalid transaction_status"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Notification Success", notificationPayload))
}
