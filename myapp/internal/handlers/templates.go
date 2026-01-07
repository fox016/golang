package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templatesDir = filepath.Join("internal", "templates")

// renderTemplate combines layout.html with a page-specific template
func renderTemplate(w http.ResponseWriter, page string, data any) {
	files := []string{
		filepath.Join(templatesDir, "layout.html"),
		filepath.Join(templatesDir, page),
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
