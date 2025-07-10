package session

import (
	"errors"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// 1. declare session errors
// ErrSessionLoad – when loading a session fails
var ErrSessionLoad = errors.New("session: load failed")

// ErrSessionSave – when saving session fails
var ErrSessionSave = errors.New("session: save failed")

// 2. Session default name
var defaultName = "session"

// 3. creates a session with MaxAge and the passed key-value
func Create(c echo.Context, key, value string) (*sessions.Session, error){
	sess, loadErr := session.Get(defaultName, c)
	if loadErr != nil {
		return nil, loadErr
	} else {		
		sess.Options = &sessions.Options{
			Path: "/",
			MaxAge: 86400,
			HttpOnly: true,
		}
		sess.Values[key] = value	
		return sess, nil
	}
}

// Set stores a key/value pair in the session and saves it.
func Set(c echo.Context, key, value string ) error {
	//? should this setter check if there is already a session with this key?
	sess, createErr := Create(c, key, value)
	if createErr != nil {
		return createErr
	} else {
		return sess.Save(c.Request(), c.Response())
	}
}

// Get retrieve a value that matches key
func GetValue(c echo.Context, key string) interface{} {
    sess, _ := session.Get(defaultName, c)
    return sess.Values[key]
}

// Delete a single key
func DeleteKey(c echo.Context, key string) error {

	sess, sessErr := session.Get(defaultName, c)
	if sessErr != nil {
		return sessErr
	} else {
		delete(sess.Values, key)
		return sess.Save(c.Request(), c.Response())
	}
}

// ClearAll wipes out the session cookie entirely.
func ClearAll(c echo.Context) error {
  sess, err := session.Get(defaultName, c)
  if err != nil {
      return err
  } else {
		sess.Options.MaxAge = -1
		return sess.Save(c.Request(), c.Response())
	}
}