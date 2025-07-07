package handler

import (
	"net/http"

	"echo-server/internal/session"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	sess session.Session
}

func NewHomeHandler(sess *session.Session) *HomeHandler {
	return &HomeHandler{sess: *sess}
}

func (hh *HomeHandler) ViewHome(c echo.Context) error {   
	if !hh.sess.Has(c.Request()) {
		return c.Redirect(http.StatusFound, "/login") 
	}

	data := map[string]bool{"Logged": true}

	return c.Render(http.StatusOK, "home.tpl", data)
}
