package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	usersdto "goGuestThePlace/dto/user"

	dto "goGuestThePlace/dto/results"

	"github.com/labstack/echo/v4"
	"goGuestThePlace/helpers"
	"goGuestThePlace/models"
	"goGuestThePlace/services"
)

type controllerUsers struct {
	UserRepository services.UserService
}

func UserController(userService services.UserService) *controllerUsers {
	return &controllerUsers{userService}
}

func (cu *controllerUsers) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := cu.UserRepository.GetUser(id)
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Users Success", user))
}

func (cu *controllerUsers) CreateUser(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)
	cu.UserRepository.CreateUser(user)
	return c.JSON(http.StatusOK, helpers.SuccessResponse("Create User Success", user))
}

func (cu *controllerUsers) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := cu.UserRepository.GetUser(id)

	request := usersdto.UpdateUserRequest{
		Fullname: c.FormValue("fullname"),
		Username: c.FormValue("username"),
		Email:    c.FormValue("email"),
	}

	if request.Fullname != "" {
		users.Fullname = request.Fullname
	}

	if request.Username != "" {
		users.Username = request.Username
	}

	if request.Email != "" {
		users.Email = request.Email
	}

	users.Updated_at = time.Now()
	users.Created_at = time.Now()

	response, err := cu.UserRepository.UpdateUser(users)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}
