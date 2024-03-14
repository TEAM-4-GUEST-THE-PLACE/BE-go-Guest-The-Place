package controller

import (
	"fmt"
	"net/http"

	"goGuestThePlace/helpers"
	"goGuestThePlace/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)
	
var DB *gorm.DB

type controllerQuestion struct{
	QuestionRepository services.QuestionService
}

func QuestionController(questionService services.QuestionService) *controllerQuestion {
	return &controllerQuestion{questionService}
}

func (cr *controllerQuestion) GetQuestions(c echo.Context) error {
	question, err := cr.QuestionRepository.GetQuestion()
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Questions Success", question))
}