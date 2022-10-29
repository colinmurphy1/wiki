package routes

import (
	"log"
	"net/http"

	"github.com/colinmurphy1/wiki/lib"
	"github.com/colinmurphy1/wiki/page"
	"github.com/colinmurphy1/wiki/state"
	"github.com/julienschmidt/httprouter"
)

func RenderPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("page")

	// Open page
	req := page.Page{
		DocumentRoot: state.Conf.Files.DocumentRoot,
		IndexPage:    state.Conf.Wiki.IndexPage,
		Path:         path,
	}
	err := req.ReadPage()

	if err != nil {
		switch err {
		case page.ErrPageNotFound:
			log.Printf("%s: Page not found", req.Path)
			lib.HTMLPage(w, http.StatusNotFound, "Not Found\n")
			return

		case page.ErrPageOutsideRoot:
			lib.HTMLPage(w, http.StatusBadRequest, "Bad Request\n")
			return

		default:
			log.Printf("%s: Error loading page: %s", req.Path, err)
			lib.HTMLPage(w, http.StatusInternalServerError, "500 Internal Server Error")
			return
		}
	}

	// Render the page
	err = req.RenderPage(w, 200)
	if err != nil {
		log.Printf("Error rendering page: %s", err)
		lib.HTMLPage(w, http.StatusInternalServerError, "Error parsing page, see the server log for details")
	}
}
