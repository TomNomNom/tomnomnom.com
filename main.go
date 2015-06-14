package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/russross/blackfriday"
)

type Blog struct {
	DevMode  bool
	Template *template.Template
	Posts    map[string]Page
}

type Page struct {
	Title   string
	Body    template.HTML
	DevMode bool
}

func (b *Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	postName := ""
	if r.URL.Path == "/" {
		postName = "index"
	} else {
		postName = r.URL.Path[7:]
	}

	post, ok := b.Posts[postName]
	if !ok {
		// Check for old-style links
		postName = strings.Replace(postName, "_", "-", -1)
		if _, ok := b.Posts[postName]; ok {
			http.Redirect(w, r, fmt.Sprintf("/posts/%s", postName), http.StatusMovedPermanently)
		} else {
			http.Error(w, "404", http.StatusNotFound)
		}
		return
	}

	err := b.Template.Execute(w, post)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	log.Println("Loading template...")
	t, err := template.ParseFiles("templates/main.html")
	if err != nil {
		log.Fatal("Failed to load template file", err)
	}

	log.Println("Loading posts...")
	dirs, err := ioutil.ReadDir("posts")
	if err != nil {
		log.Fatal("Failed to read posts", err)
	}

	posts := make(map[string]Page)
	for _, dir := range dirs {
		b, err := ioutil.ReadFile(fmt.Sprintf("posts/%s/article.mkd", dir.Name()))
		if err != nil {
			continue
		}
		fistLine := strings.SplitN(string(b), "\n", 2)[0]
		posts[dir.Name()] = Page{
			Title:   fmt.Sprintf("%s - TomNomNom.com", strings.Trim(fistLine, "# ")),
			Body:    template.HTML(blackfriday.MarkdownCommon(b)),
			DevMode: true,
		}
	}

	b := &Blog{Template: t, Posts: posts}
	http.Handle("/", b)

	http.Handle("/styles/", http.FileServer(http.Dir("public")))
	http.Handle("/images/", http.FileServer(http.Dir("public")))
	http.Handle("/googlebfd35bb0eb2d4f45.html", http.FileServer(http.Dir("public")))
	http.Handle("/keybase.txt", http.FileServer(http.Dir("public")))

	log.Println("Starting up...")
	http.ListenAndServe(":9090", nil)
}
