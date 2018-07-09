package main

import (
	"../src/web"
	"../src/utils"
	"../src/dto"
	"../src/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func SecurityContextPasser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Request().Cookie(config.X_AUTH_TOKEN_COOKIE_NAME)
		if err != nil {
			return c.JSON(http.StatusOK, dto.Error("Not authorized"))
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWT_SALT), nil
		})

		if token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				c.Set("login", claims["login"])
			}
			return next(c)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			utils.RemoveCookie(c)
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return c.JSON(http.StatusOK, dto.Error("Not valid token"))
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return c.JSON(http.StatusOK, dto.Error("Outdated token"))
			} else {
				return c.JSON(http.StatusOK, dto.Error("Invalid token"))
			}
		} else {
			utils.RemoveCookie(c)
			return c.JSON(http.StatusOK, dto.Error("Not authorized"))
		}
	}
}

func main() {
	router := echo.New()

	router.Use(middleware.Logger())

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	securityRouterGroup := router.Group("/security")

	securityRouterGroup.POST("/register", web.Reg)
	securityRouterGroup.POST("/login", web.Login)
	securityRouterGroup.GET("/logout", web.Logout)

	securedRoutesInSecurityGroup := router.Group("/security/principal")
	securedRoutesInSecurityGroup.Use(SecurityContextPasser)
	securedRoutesInSecurityGroup.GET("", web.Principal)

	subscriptionRouterGroup := router.Group("/subscription")
	subscriptionRouterGroup.Use(SecurityContextPasser)

	subscriptionRouterGroup.GET("/list", web.GetSubscriptions)
	subscriptionRouterGroup.POST("/add", web.AddSubscription)
	subscriptionRouterGroup.POST("/remove", web.RemoveSubscription)
	subscriptionRouterGroup.POST("/edit", web.EditSubscription)

	router.GET("/feed", web.GetFeed)

	router.Start(":8088")
}
