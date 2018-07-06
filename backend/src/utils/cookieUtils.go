package utils

import (
	"github.com/labstack/echo"
	"../../src/config"
	"net/http"
	"time"
)

func SetXAuthCookieValue(context echo.Context, token string) echo.Context {
	cookie := new(http.Cookie)
	cookie.Name = config.X_AUTH_TOKEN_COOKIE_NAME
	cookie.Value = token
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	context.SetCookie(cookie)
	return context
}

func RemoveCookie(context echo.Context) echo.Context {
	cookie := new(http.Cookie)
	cookie.Name = config.X_AUTH_TOKEN_COOKIE_NAME
	cookie.Path = "/"
	cookie.Value = ""
	cookie.Expires = time.Unix(0, 0)
	context.SetCookie(cookie)
	return context
}