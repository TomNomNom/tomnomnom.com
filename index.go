package main

import (
	"html/template"
	"log"
	"net/http"
)

type IndexHandler struct {
	DevMode  bool
	Template *template.Template
	Post     *Page
}

func (i *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	err := i.Template.Execute(w, i.Post)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}
}
