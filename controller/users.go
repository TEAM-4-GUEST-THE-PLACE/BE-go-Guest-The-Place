package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	// "strconv"
	// "time"

	// usersdto "goGuestThePlace/dto/user"

	// dto "goGuestThePlace/dto/results"

	"goGuestThePlace/helpers"
	"goGuestThePlace/models"

	"github.com/labstack/echo/v4"

	// "goGuestThePlace/models"
	"goGuestThePlace/services"
)

type controllerUsers struct {
	UserRepository services.UserService
}

func UserController(userService services.UserService) *controllerUsers {
	return &controllerUsers{userService}
}

func (cu *controllerUsers) GetUser(c echo.Context) error {
	user, err := cu.UserRepository.GetUser()
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Users Success", user))
}

func (cu *controllerUsers) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := cu.UserRepository.UpadateUser(id, models.Users{})

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	request := models.UpdateUserRequest{
		Fullname: c.FormValue("fullname"),
		Username: c.FormValue("username"),
		Email:    c.FormValue("email"),
	}

	if request.Fullname == "" {
		request.Fullname = users.Fullname
	}

	if request.Username == "" {
		request.Username = users.Username
	}

	if request.Email == "" {
		request.Email = users.Email
	}

	users.Updated_at = time.Now()

	response, err := cu.UserRepository.UpadateUser(id, users)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Update User Success", response))
}
