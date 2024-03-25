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

// func (cr *controllerTransaction) Notification(c echo.Context) error {
// 	var notificationPayload map[string]interface{}

// 	if err := c.Bind(&notificationPayload); err != nil {
// 		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
// 	}
// 	transactionStatus := notificationPayload["transaction_status"].(string)
// 	fraudStatus := notificationPayload["fraud_status"].(string)
// 	orderId := notificationPayload["order_id"].(string)

// 	orderId, _ = strconv.Atoi(orderId)

// 	fmt.Print("test payload", notificationPayload)

// 	if transactionStatus == "capture" {
// 		if fraudStatus == "accept" {
// 			TransactionRepository.UpadateTransaction("success", orderId)

// 		} else if fraudStatus == "challenge" {
// 			TransactionRepository.UpadateTransaction("pending", orderId)

// 		}
// 	} else if transactionStatus == "settlement" {
// 		TransactionRepository.UpadateTransaction("success", orderId)

// 	} else if transactionStatus == "deny" {
// 		TransactionRepository.UpadateTransaction("failed", orderId)

// 	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
// 		TransactionRepository.UpadateTransaction("failed", orderId)

// 	} else if transactionStatus == "pending" {
// 		tTransactionRepository, err := UpadateTransaction("pending", orderId)

// 	} else if transactionStatus == "settlement" {
// 		TransactionRepository, err := UpadateTransaction("success", orderId)

// 	}

// 	return c.JSON(http.StatusOK, helpers.SuccessResponse("Notification Success", transaction))

// }
