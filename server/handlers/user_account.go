package handlers

import (
	"net/http"
)

// type UserAccountHandle struct {
// }

func Register(w http.ResponseWriter, r *http.Request) {
	// logger := hlog.FromRequest(r)

	// check email is not being used
	// validate password
	// create user in DB
	// return token
}

func Login(w http.ResponseWriter, r *http.Request) {
	// logger := hlog.FromRequest(r)

	// need to research
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// logger := hlog.FromRequest(r)

	// need to research
}
