package web

import (
	"github.com/labstack/echo"
	"net/http"
)

func AddSubscription(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}

func GetSubscriptions(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}

func RemoveSubscription(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}

func EditSubscription(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}
