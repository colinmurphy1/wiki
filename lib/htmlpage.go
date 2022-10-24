package lib

import (
	"fmt"
	"net/http"
)

func HTMLPage(w http.ResponseWriter, httpCode int, content interface{}) {
	// Set content-type to text/html with an implicit utf-8 charset
	w.Header().Set("Content-type", "text/html;charset=utf-8")
	w.WriteHeader(httpCode)
	fmt.Fprint(w, content)
}
