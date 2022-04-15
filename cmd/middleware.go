package main

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
)

var session *scs.SessionManager

// SessionLoad loads the session on requests
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}