package handler

import (
	"echo-server/internal/service"
	"echo-server/internal/session"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userSvc service.UserService
	session session.Session
}

func NewAuthHandler(userService service.UserService, sess *session.Session) *AuthHandler{
	return &AuthHandler{userSvc: userService, session: *sess}
}

// Handlers

// GET /register
func (ah *AuthHandler) RegisterForm(c echo.Context) error{
	// if !ah.session.Has(c.Request()) {
	// 	return c.Redirect(http.StatusSeeOther, "/login")
	// }
	return c.Render(http.StatusOK, "register.tpl", nil)
}

// POST /register
func (ah *AuthHandler) RegisterSubmit(c echo.Context) error {
	usr:= c.FormValue("username")
	email := c.FormValue("email")
	pwd := c.FormValue("password")
	rep := c.FormValue("repeatedPassword")

	if usr == "" || email == "" || pwd == "" || pwd != rep {
		data := map[string]string{"Error": "Fill all fields and ensure passwords match"}
		return c.Render(http.StatusOK, "register.tpl", data)
	}

	_, registerErr := ah.userSvc.Register(usr, email, pwd)
	if registerErr != nil {
		if errors.Is(registerErr, service.ErrUserExist) {
			return c.Render(http.StatusOK, "register.tpl", map[string]string{"Error": registerErr.Error()})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "server error")
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}

// GET /login
func (ah *AuthHandler) LoginForm(c echo.Context) error {
	// if !ah.session.Has(c.Request()) {
	// 	return c.Redirect(http.StatusSeeOther, "/login")
	// }
	return c.Render(http.StatusOK, "login.tpl", nil)
}

// POST /login
func (ah *AuthHandler) LoginSubmit(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, loginErr := ah.userSvc.Login(username, password)
	if loginErr != nil {
		code := http.StatusInternalServerError
		if errors.Is(loginErr, service.ErrInvalidCredentials) {
			code = http.StatusUnauthorized
		}
		return echo.NewHTTPError(code, loginErr.Error())
	}

	sessErr := ah.session.Set(c.Response().Writer, c.Request(), "user_id", strconv.Itoa(user.ID))
	if  sessErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "session error")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

// GET /logout
func (ah *AuthHandler) Logout(c echo.Context) error {
	sessErr := ah.session.Delete(c.Response().Writer, c.Request())
	if sessErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, sessErr.Error())
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}