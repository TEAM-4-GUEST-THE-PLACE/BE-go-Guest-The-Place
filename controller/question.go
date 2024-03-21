package controller

import (
	"net/http"
	"strconv"

	"goGuestThePlace/helpers"
	"goGuestThePlace/services"

	"github.com/labstack/echo/v4"
)

type controllerQuestion struct {
	QuestionRepository services.QuestionService
}

func QuestionController(questionService services.QuestionService) *controllerQuestion {
	return &controllerQuestion{questionService}
}

func (cr *controllerQuestion) GetQuestions(c echo.Context) error {
	question, err := cr.QuestionRepository.GetQuestion()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Questions Success", question))
}

func (cr *controllerQuestion) GetQuestionById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	question, err := cr.QuestionRepository.GetQuestionById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	// err = json.Unmarshal([]byte(question.Options), &question)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	// }

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Question Success with id", question))

}
