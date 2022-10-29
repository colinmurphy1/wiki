package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/colinmurphy1/wiki/lib"
	"github.com/colinmurphy1/wiki/state"
	"github.com/julienschmidt/httprouter"
)

func LoginForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lib.HTMLPage(w, http.StatusOK, "Login Form")
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Ensure username and password are specified
	if len(username) == 0 || len(password) == 0 {
		// Redirect back to login page
		lib.Redirect(w, false, "/login")
		return
	}

	// Attempt login
	u, err := state.Users.Authenticate(username, password)
	if err != nil {
		lib.HTMLPage(w, http.StatusUnauthorized, "Could not sign in\n")
		log.Printf("auth: %s failed to log in: %s", username, err)
		return
	}

	lib.HTMLPage(w, http.StatusOK, fmt.Sprintf("Hello, %s!\n", u.Username))
	log.Printf("auth: %s logged in successfully", username)
}
