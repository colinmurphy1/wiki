package page

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Lookup order. If two files with the same name exist, the first one that matches will be rendered
// TODO: Make this be configurable in the config

var (
	FILENAME_ORDER     = []string{"md", "html", "txt"}
	ErrPageOutsideRoot = errors.New("page is outside of document root")
	ErrPageNotFound    = errors.New("page not found")
)

// Page struct
type Page struct {
	DocumentRoot string // Document root
	Path         string // path to the page, including full file name
	IndexPage    string // The name of the default page to see when going to a namespace
	Namespace    string // Page Namespace
	FileName     string // Page file name (without the format)
	ContentType  string // Page Content type
	Content      []byte // Page contents
}

// Opens a page and returns the contents as a string, or an error if there is one
// This will follow all pathing and filename lookup orders
func (p *Page) ReadPage() error {
	// Identify full path to document root
	documentRoot, _ := filepath.Abs(p.DocumentRoot)

	// If no path is specified (or you go to /p/), set the main namespace and index page as the page to view
	if len(p.Path) == 0 || p.Path == "/" {
		p.Path = p.IndexPage
		p.FileName = p.IndexPage
		p.Namespace = ""
	} else {
		// Find namespace and page to view
		pathMatch := regexp.MustCompile("^(.*/)([^/]*)$").FindStringSubmatch(p.Path)

		p.Namespace = pathMatch[1]
		p.FileName = pathMatch[2]

		// If no page is specified, use the index page
		if pathMatch[2] == "" {
			p.Path = p.Path + p.IndexPage
		}
	}

	// FULL path to the page being viewed
	absPath := filepath.Clean(documentRoot + "/" + p.Path)

	// Identify directory traversal attacks (https://owasp.org/www-community/attacks/Path_Traversal)
	if !strings.HasPrefix(absPath, documentRoot) {
		return ErrPageOutsideRoot
	}

	// Open file using filename order
	matchFound := false
	for _, ext := range FILENAME_ORDER {
		file, err := os.ReadFile(absPath + "." + ext)
		// Try to find a file match
		if err == nil {
			matchFound = true
			p.ContentType = ext
			p.Content = file
			break
		}
	}

	// Return error if no match is found
	if !matchFound {
		return ErrPageNotFound
	}

	return nil
}
