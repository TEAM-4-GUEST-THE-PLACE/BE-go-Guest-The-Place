package controller

import (
	"goGuestThePlace/helpers"
	"goGuestThePlace/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"goGuestThePlace/models"
)

type controllerAuth struct {
	AuthService services.AuthService
}

func AuthController(authService services.AuthService) *controllerAuth {
	return &controllerAuth{authService}
}

func (ca *controllerAuth) Login(c echo.Context) error {
	request := new(models.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	user := models.Users{
		Fullname: request.Fullname,
		Email:    request.Email,
	}

	user, err := ca.AuthService.Login(user.Fullname)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	LoginResponse := models.LoginResponse{
		Fullname:       user.Fullname,
		Email:          user.Email,
		Diamonds_totals: user.Diamonds_totals,
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Login Success", LoginResponse))
}
