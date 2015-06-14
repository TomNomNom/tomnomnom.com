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
	Posts    map[string][]byte
}

type Page struct {
	Title   string
	Body    template.HTML
	DevMode bool
}

func (b *Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	post := ""
	if r.URL.Path == "/" {
		post = "index"
	} else {
		post = r.URL.Path[7:]
	}

	raw, ok := b.Posts[post]
	if !ok {
		// Check for old-style links
		post = strings.Replace(post, "_", "-", -1)
		if _, ok := b.Posts[post]; ok {
			http.Redirect(w, r, fmt.Sprintf("/posts/%s", post), http.StatusMovedPermanently)
		} else {
			http.Error(w, "404", http.StatusNotFound)
		}
		return
	}

	html := blackfriday.MarkdownCommon(raw)

	b.Template.New("body").Parse(string(html))

	p := &Page{
		Title:   "Tom Hudson's blog - TomNomNom.com",
		Body:    template.HTML(html),
		DevMode: b.DevMode,
	}
	err := b.Template.Execute(w, p)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	t, err := template.ParseFiles("templates/main.html")
	if err != nil {
		log.Fatal("Failed to load template file", err)
	}

	dirs, err := ioutil.ReadDir("posts")
	if err != nil {
		log.Fatal("Failed to read posts", err)
	}
	posts := make(map[string][]byte)
	for _, dir := range dirs {
		b, err := ioutil.ReadFile(fmt.Sprintf("posts/%s/article.mkd", dir.Name()))
		if err != nil {
			continue
		}
		posts[dir.Name()] = b
	}

	b := &Blog{DevMode: true, Template: t, Posts: posts}
	http.Handle("/", b)

	http.Handle("/styles/", http.FileServer(http.Dir("public")))
	http.Handle("/images/", http.FileServer(http.Dir("public")))
	http.Handle("/googlebfd35bb0eb2d4f45.html", http.FileServer(http.Dir("public")))
	http.Handle("/keybase.txt", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":9090", nil)
}
