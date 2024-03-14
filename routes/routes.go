package routes

import (

	"goGuestThePlace/config"
	"goGuestThePlace/controller"
	"goGuestThePlace/services"

	"github.com/labstack/echo"
)

func Routes(e *echo.Group) {

	ra := services.RepositoryQuestion(config.DB)
	cra := controller.QuestionController(ra)

	e.GET("/Questions", cra.GetQuestions)

}
