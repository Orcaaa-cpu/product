package itemsproductcontroller

import (
	"net/http"
	"product/helper"
	"text/template"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	helper.Lock.Lock()
	defer helper.Lock.Unlock()

	sess, _ := session.Get("session", c)
	if len(sess.Values) == 0 || sess.Values["loggedIn"] != true {
		c.Redirect(http.StatusSeeOther, "/login")
	}

	temp, err := template.ParseFiles("view/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"name": sess.Values["name"],
	}

	temp.Execute(c.Response().Writer, data)

	return c.NoContent(http.StatusOK)
}
