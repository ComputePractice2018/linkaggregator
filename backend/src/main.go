package main

import (
	"../src/web"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	router := echo.New()

	router.Use(middleware.Logger())

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	securityRouterGroup := router.Group("/security")

	securityRouterGroup.GET("", web.Principal)
	securityRouterGroup.POST("/register", web.Reg)
	securityRouterGroup.POST("/login", web.Login)
	securityRouterGroup.GET("/logout", web.Logout)

	subscriptionRouterGroup := router.Group("/subscription")

	subscriptionRouterGroup.GET("", web.GetSubscriptions)
	subscriptionRouterGroup.POST("/add", web.AddSubscription)
	subscriptionRouterGroup.POST("/remove", web.RemoveSubscription)
	subscriptionRouterGroup.POST("/edit", web.EditSubscription)

	router.GET("/feed", web.GetFeed)

	router.Start(":8088")
}
