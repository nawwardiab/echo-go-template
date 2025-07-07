package session

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// 1. declare session errors
// ErrSessionLoad – when loading a session fails
var ErrSessionLoad = errors.New("session: load failed")

// ErrSessionSave – when saving session fails
var ErrSessionSave = errors.New("session: save failed")

// 2. Session struct type
type Session struct {
	store sessions.Store
	name string
	options *sessions.Options
}

// 3. create NewSession func
// NewSession manages creating a session using provided secret key
func NewSession(key string) *Session {
	// 3.1 create store
	store := sessions.NewCookieStore([]byte(key))

	// 3.2 configure session's options
	opts := &sessions.Options{
		Path: "/",
		MaxAge: 86400,
		HttpOnly: true,
	}
	// 3.3 inst session from the Session struct
	session := &Session{
		store: store,
		name: "session",
		options: opts,
	}

	// 3.4 return instance
	return session
}

// load retrieves the session or returns ErrSessionLoad.
func (s *Session) load(r *http.Request) (*sessions.Session, error) {
	sess, sessErr := s.store.Get(r, s.name)
	if sessErr != nil {
		return nil, sessErr
	}

	sess.Options = s.options
	return sess, nil
}

// Set stores a key/value pair in the session and saves it.
func (s *Session) Set(w http.ResponseWriter, r *http.Request, key, value string) error {
	sess, loadErr := s.load(r)
	if loadErr != nil {
		return loadErr
	}

	sess.Values[key] = value
	saveErr := sess.Save(r, w)
	if saveErr != nil {
		return fmt.Errorf("%w: %v", ErrSessionSave, saveErr)
	} else {
		return nil
	}

}

// Get retrieve raw value that matches key
func (s *Session) Get(r *http.Request, key string) (string, error) {
    sess, sessErr := s.store.Get(r, s.name)
    if sessErr != nil {
        return "", sessErr
    }
    rawValue, ok := sess.Values[key]
    if !ok {
        return "", nil // not set
    }
    str, ok := rawValue.(string)
    if !ok {
        return "", fmt.Errorf("session: value for %q is not a string", key)
    }
    return str, nil
}

// Has returns true if a non-new session exists.
func (s *Session) Has(r *http.Request) bool {

	sess, sessErr := s.load(r)
	bl := sessErr == nil && !sess.IsNew 

	return bl 
}

// Delete invalidates the session cookie.
func (s *Session) Delete(w http.ResponseWriter, r *http.Request) error {

	sess, sessErr := s.load(r)
	if sessErr != nil {
		http.Error(w, sessErr.Error(), http.StatusInternalServerError)
	}

	sess.Options = &sessions.Options{
		Path: "/",
		MaxAge: -1,
		HttpOnly: true,
	}

	sessErr = sess.Save(r, w)
	if sessErr != nil {
		http.Error(w, sessErr.Error(), http.StatusInternalServerError)
	}

	return sessErr
}