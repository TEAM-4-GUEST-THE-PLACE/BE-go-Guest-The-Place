package controller

import (
	"fmt"
	"net/http"

	"goGuestThePlace/helpers"
	"goGuestThePlace/services"

	"github.com/labstack/echo/v4"

)



type controllerDiamond struct{
	DiamondRepository services.DiamondService
}


func DiamondController(diamondService services.DiamondService) *controllerDiamond {
	return &controllerDiamond{diamondService}
}

func (cd *controllerDiamond) GetDiamond(c echo.Context) error {
	diamond, err := cd.DiamondRepository.GetDiamond()
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}


	return c.JSON(http.StatusOK, helpers.SuccessResponse("Get Diamonds Success", diamond))
}