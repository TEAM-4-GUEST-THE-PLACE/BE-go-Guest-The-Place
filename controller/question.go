package controller

import (
	"net/http"

	"goGuestThePlace/services"
	"goGuestThePlace/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)
	
var DB *gorm.DB

type controllerQuestion struct{
	QuestionRepository services.QuestionService
}

func QuestionController(articleServices services.QuestionService) *controllerQuestion {
	return &controllerQuestion{articleServices}
}

func (cr *controllerQuestion) GetQuestions(c echo.Context) error {
	question, err := cr.QuestionRepository.GetQuestion()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get 	Questions Success", question))
}