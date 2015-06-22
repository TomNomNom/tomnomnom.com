package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title   string
	Body    template.HTML
	DevMode bool
}

type PageError struct {
	Message  string
	Code     int
	Location string
}

func (e *PageError) Error() string {
	return e.Message
}

func main() {
	log.Println("Loading template...")
	t, err := template.ParseFiles("templates/main.html")
	if err != nil {
		log.Fatal("Failed to load template file", err)
	}

	log.Println("Loading posts...")
	posts := LoadPosts()

	i := &IndexHandler{Template: t, Post: posts["index"]}
	http.Handle("/", i)

	p := &PostsHandler{Template: t, Posts: posts}
	http.Handle("/posts/", p)

	http.Handle("/styles/", http.FileServer(http.Dir("public")))
	http.Handle("/images/", http.FileServer(http.Dir("public")))
	http.Handle("/googlebfd35bb0eb2d4f45.html", http.FileServer(http.Dir("public")))
	http.Handle("/keybase.txt", http.FileServer(http.Dir("public")))

	log.Println("Starting up...")
	http.ListenAndServe(":9090", nil)
}
