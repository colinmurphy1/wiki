package lib

import "net/http"

/*
Redirects the user a specified address
If permanent is true, the redirect will be permanent
*/
func Redirect(w http.ResponseWriter, permanent bool, address string) {
	httpCode := http.StatusTemporaryRedirect
	if permanent {
		httpCode = http.StatusPermanentRedirect
	}
	w.WriteHeader(httpCode)
	w.Header().Set("Location", address)
}
