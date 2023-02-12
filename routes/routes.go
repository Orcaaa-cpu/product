package routes

import (
	autproductcontroller "product/aut-product/aut-product-controller"
	itemsproductcontroller "product/items-product/items-product-controller"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/login", autproductcontroller.ViewLogin)
	e.POST("/login", autproductcontroller.Login)
	e.GET("/logout", autproductcontroller.Logout)

	e.GET("/index", itemsproductcontroller.Index)

	return e
}
