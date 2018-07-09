package web

import (
	"github.com/labstack/echo"
	"net/http"
	"../../src/dto"
)

func AddSubscription(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}

func GetSubscriptions(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.Success())
}

func RemoveSubscription(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}

func EditSubscription(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}
