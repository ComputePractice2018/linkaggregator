package web

import (
	"github.com/labstack/echo"
	"../../src/dto"
	"../../src/utils"
	"../../src/service"
	"net/http"
	"fmt"
)

func AddSubscription(requestContext echo.Context) error {
	url := requestContext.FormValue("url")

	if utils.IsEmpty(url) {
		return requestContext.JSON(http.StatusOK, dto.Error("Empty URL"))
	}

	login := requestContext.Get("login")
	textLogin := fmt.Sprintf("%v", login)

	service.SubscribeUserToRSS(textLogin, url)

	return requestContext.JSON(http.StatusOK, dto.Success())
}

func GetSubscriptions(requestContext echo.Context) error { /*
	login := requestContext.Get("login")
	textLogin := fmt.Sprintf("%v", login)

	res := service.GetUserSubscriptions(textLogin)*/

	return requestContext.JSON(http.StatusOK, dto.Success())
}

func RemoveSubscription(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}

func EditSubscription(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}
