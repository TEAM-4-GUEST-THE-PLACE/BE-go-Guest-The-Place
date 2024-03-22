package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"goGuestThePlace/helpers"
	"goGuestThePlace/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

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

func (cu *controllerUsers) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}
	user, err := cu.UserRepository.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Users Success with id", user))
}

func (cu *controllerUsers) GetUserByEmail(c echo.Context) error {
	email := c.Param("email")
	user, err := cu.UserRepository.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Users Success with email", user))

}



func (cu *controllerUsers) CreatedUser(c echo.Context) error {
	request := new(models.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	data := models.Users{
		ID:             request.ID,
		Avatar_id:      request.Avatar_id,
		Diamond_totals: request.Diamond_totals,
		Fullname:       request.Fullname,
		Email:          request.Email,

		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	response, err := cu.UserRepository.CreatedUser(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Create User Success", response))

}

func (cu *controllerUsers) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := cu.UserRepository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	request := new(models.UpdateUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))

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

	response, err := cu.UserRepository.UpadateUser(users)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Update User Success", response))
}

func (cu *controllerUsers) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	users, err := cu.UserRepository.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	data, err := cu.UserRepository.DeleteUsers(users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Delete User Success", data))
}
