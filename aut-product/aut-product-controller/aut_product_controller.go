package autproductcontroller

import (
	"errors"
	"html/template"
	"net/http"
	autproductmodel "product/aut-product/aut-product-model"
	"product/entities"
	"product/helper"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func ViewLogin(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	temp, err := template.ParseFiles("view/login.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	temp.Execute(c.Response().Writer, nil)
	return c.NoContent(http.StatusOK)
}

func Login(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	username := c.FormValue("username")
	password := c.FormValue("password")

	user := entities.Users{}

	err := autproductmodel.Login(&user, username, password)
	if err != nil {
		err = errors.New("Username atau Password salah")
		data := map[string]interface{}{
			"error": err,
		}
		temp, _ := template.ParseFiles("view/login.html")
		temp.Execute(c.Response().Writer, data)
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["loggedIn"] = true
	sess.Values["email"] = user.Email
	sess.Values["username"] = user.Username
	sess.Values["name"] = user.Name

	sess.Save(c.Request(), c.Response())

	c.Redirect(http.StatusSeeOther, "/index")

	return c.NoContent(http.StatusOK)
}

func Logout(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	sess, _ := session.Get("session", c)

	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	c.Redirect(http.StatusSeeOther, "/login")

	return c.NoContent(http.StatusOK)
}
