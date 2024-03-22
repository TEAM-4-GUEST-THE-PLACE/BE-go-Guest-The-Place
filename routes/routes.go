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
	e.GET("/questions/:id", cra.GetQuestionById)

	rpx := services.RepositoryUser(config.DB)
	crpa := controller.UserController(rpx)	

	e.GET("/users", crpa.GetUser)
	e.GET("/users/:id", crpa.GetUserById)
	e.POST("/users", crpa.CreatedUser)
	e.PUT("/users/:id", crpa.UpdateUser)
	e.DELETE("/users/:id", crpa.DeleteUser)
	e.GET("/users/:email", crpa.GetUserByEmail)

	rpa := services.RepositoryDiamond(config.DB)
	crdpa := controller.DiamondController(rpa)

	e.GET("/diamonds", crdpa.GetDiamond)

	ca := services.RepositoryAvatar(config.DB)
	craa := controller.AvatarController(ca)

	e.GET("/avatars", craa.GetAvatar)

	cr := services.RepositoryTransaction(config.DB)
	crt := controller.TransactionController(cr)

	e.GET("/transactions", crt.GetTransactions)
	e.POST("/transactions", crt.CreateTransaction)

	crty := services.RepositoryAuth(config.DB)
	ct := controller.AuthController(crty)

	e.POST("/login", ct.Login)

}
