package web

import (
	"../../src/utils"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func Reg(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}

func Login(c echo.Context) error {

	login := c.FormValue("login")
	password := c.FormValue("password")

	if utils.IsEmpty(login) || utils.IsEmpty(password) {
		panic("Empty credentials")
	}

	fmt.Print(login, password)
	return c.String(http.StatusOK, "{}")
}

func Principal(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}

func Logout(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}

func throwError(ctx echo.Context, cause string) {
}
