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
}

func NewAuthHandler(userService service.UserService) *AuthHandler{
	return &AuthHandler{userSvc: userService}
}

// Handlers

// GET /register
func (ah *AuthHandler) RegisterForm(c echo.Context) error {
	if session.GetValue(c, "user_id") != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	} else {
		return c.Render(http.StatusOK, "register.tpl", nil)
	}	
}

// POST /register
func (ah *AuthHandler) RegisterSubmit(c echo.Context) error {
	usr:= c.FormValue("username")
	email := c.FormValue("email")
	pwd := c.FormValue("password")
	rep := c.FormValue("repeatedPassword")
	
	if usr == "" || email == "" || pwd == "" {
		data := map[string]string{"Error": "Fill all fields"}
		return c.Render(http.StatusOK, "register.tpl", data)
	} else if pwd != rep {
		data := map[string]string{"Error": "Password and repeated password must match"}			
		return c.Render(http.StatusOK, "register.tpl", data)
	} 
		
	_, registerErr := ah.userSvc.Register(usr, email, pwd)
	if	registerErr != nil {
		if errors.Is(registerErr, service.ErrUserExist) {
			return c.Render(http.StatusOK, "register.tpl", map[string]string{"Error": registerErr.Error()})
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, "server error")
		}
	} else {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
}

// GET /login
func (ah *AuthHandler) LoginForm(c echo.Context) error {
	if  session.GetValue(c, "user_id") != nil {
		return c.Redirect(http.StatusSeeOther, "/")
	} else {
		return c.Render(http.StatusOK, "login.tpl", nil)
	}
}

// POST /login
func (ah *AuthHandler) LoginSubmit(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	//TODO sanitize Form data
	user, loginErr := ah.userSvc.Login(username, password)

	if loginErr != nil {
		code := http.StatusInternalServerError
		if errors.Is(loginErr, service.ErrInvalidCredentials) {
			code = http.StatusUnauthorized
		}
		return echo.NewHTTPError(code, loginErr.Error())
	} else {	
		sessErr := session.Set(c, "user_id", strconv.Itoa(user.ID))
		if  sessErr != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "session error")
		} else {			
			return c.Redirect(http.StatusSeeOther, "/")
		}
	}
}

// GET /logout – deletes user_id from session and redirects to login
func (ah *AuthHandler) Logout(c echo.Context) error {
	sessErr := session.DeleteKey(c, "user_id")
	if sessErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, sessErr.Error())
	} else {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
}