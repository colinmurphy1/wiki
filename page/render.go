package page

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/colinmurphy1/wiki/state"
	"github.com/gomarkdown/markdown"
)

// HTML Template struct
type HTMLPage struct {
	Site    state.ConfigWiki // Site settings
	Content string           // HTML content
}

// Render a page
func (p *Page) RenderPage(w http.ResponseWriter, httpCode int) {
	var pageContent string
	switch p.ContentType {
	// Markdown
	case "md", "markdown":
		pageContent = string(markdown.ToHTML(p.Content, nil, nil))

	// Plain text
	case "txt":
		pageContent = "<pre>\n" + string(p.Content) + "\n</pre>\n"

	// HTML - TODO: Sanitize HTML output
	default:
		pageContent = string(p.Content)
	}

	// Get template file
	tmplPath, _ := filepath.Abs("./templates/" + state.Conf.Wiki.Theme + "/page.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("Error loading template: %s", err)
	}

	// Render template and write response
	w.WriteHeader(httpCode)
	tmpl.Execute(w, HTMLPage{
		Site:    state.Conf.Wiki, // pass site config over
		Content: pageContent,     // rendered html
	})
}
