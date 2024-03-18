package controller


import (
	"fmt"
	"net/http"

	"goGuestThePlace/helpers"
	"goGuestThePlace/services"


	"github.com/labstack/echo/v4"
)


type controllerAvatar struct{
	AvatarRepository services.AvatarService
}


func AvatarController(avatarService services.AvatarService) *controllerAvatar {
	return &controllerAvatar{avatarService}
}

func (ca *controllerAvatar) GetAvatar(c echo.Context) error {
	avatar, err := ca.AvatarRepository.GetAvatar()
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Avatars Success", avatar))
}