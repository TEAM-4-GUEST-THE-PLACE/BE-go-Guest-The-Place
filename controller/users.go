package controller


import (
	"fmt"
	"net/http"

	"goGuestThePlace/helpers"
	"goGuestThePlace/services"

	"github.com/labstack/echo/v4"

)



type controllerUsers struct{
	UserRepository services.UserService
}


func UserController(userService services.UserService) *controllerUsers {
	return &controllerUsers{userService}
}

func (cu *controllerUsers) GetUsers(c echo.Context) error {
	user, err := cu.UserRepository.GetUser()
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Users Success", user))
}