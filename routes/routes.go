package routes

import (
	"goGuestThePlace/config"
	"goGuestThePlace/controller"
	"goGuestThePlace/services"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group) {

	ra := services.RepositoryQuestion(config.DB)
	cra := controller.QuestionController(ra)

	e.GET("/questions", cra.GetQuestions)


	rp := services.RepositoryUser(config.DB)
	crpa := controller.UserController(rp)

	e.GET("/users", crpa.GetUsers)

}
