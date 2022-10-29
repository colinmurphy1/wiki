package page

import (
	"html"
	"net/http"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/colinmurphy1/wiki/state"
	"github.com/gomarkdown/markdown"
)

// HTML Template struct
type HTMLPage struct {
	Site      state.ConfigWiki // Site settings
	PageTitle string           // Page title
	Content   string           // HTML content
}

// Gets the page title from the h1 or h2 tags. Returns the page title, or an empty string if there is none.
func getTitleFromHTML(content string) string {
	search := regexp.MustCompile("<h[1-2]>(.+)</h[1-2]>").FindStringSubmatch(content)

	// No matches, return empty string
	if len(search) == 0 {
		return ""
	}

	// First match is the page title
	return search[1]
}

// Render a page
func (p *Page) RenderPage(w http.ResponseWriter, httpCode int) error {
	var pageContent string
	switch p.ContentType {
	// Markdown
	case "md", "markdown":
		pageContent = string(markdown.ToHTML(p.Content, nil, nil))

	// Plain text
	case "txt":
		pageContent = "<pre>\n" + html.EscapeString(string(p.Content)) + "\n</pre>\n"

	// HTML - TODO: Sanitize HTML output
	default:
		pageContent = string(p.Content)
	}

	// Find title of page. if there is no title, use the filename of the page
	pageTitle := getTitleFromHTML(pageContent)
	if pageTitle == "" {
		pageTitle = p.FileName
	}

	// Get template file
	tmplPath, _ := filepath.Abs(state.Conf.Files.ThemeDir + "/page.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	// Render template and write response
	w.WriteHeader(httpCode)
	tmpl.Execute(w, HTMLPage{
		Site:      state.Conf.Wiki, // pass site config over
		PageTitle: pageTitle,
		Content:   pageContent, // rendered html
	})

	return nil
}
