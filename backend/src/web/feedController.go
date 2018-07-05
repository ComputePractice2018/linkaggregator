package web

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetFeed(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}
