package main

import (
	"html/template"
	"log"
	"path/filepath"
)

func mergeFilepaths() []string {
	directories := []string{
		"./templates/",
		"./templates/components/",
	}

	files := []string{
		"./index.html",
	}

	templatePattern := "*.html"

	for _, dir := range directories {
		pattern := filepath.Join(dir, templatePattern)
		matchedFiles, err := filepath.Glob(pattern)
		if err != nil {
			log.Fatal(err)
		}
		files = append(files, matchedFiles...)
	}

	return files
}

func initAndBumdleTemplates() {
	files := mergeFilepaths()
	tmpl = template.Must(template.ParseFiles(files...))
}
