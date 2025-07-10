package handler

import (
	"net/http"

	"echo-server/internal/session"

	"github.com/labstack/echo/v4"
)



func ViewHome(c echo.Context) error {   
	if session.GetValue(c, "user_id") == nil {
		return c.Redirect(http.StatusFound, "/login") 
	} else {
		data := map[string]bool{"Logged": true}
		return c.Render(http.StatusOK, "home.tpl", data)
	}
}
