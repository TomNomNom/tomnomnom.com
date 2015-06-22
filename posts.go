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

func LoadPosts() map[string]*Page {
	dirs, err := ioutil.ReadDir("posts")
	if err != nil {
		log.Fatal("Failed to read posts", err)
	}

	posts := make(map[string]*Page)
	for _, dir := range dirs {
		b, err := ioutil.ReadFile(fmt.Sprintf("posts/%s/article.mkd", dir.Name()))
		if err != nil {
			continue
		}
		fistLine := strings.SplitN(string(b), "\n", 2)[0]
		posts[dir.Name()] = &Page{
			Title:   fmt.Sprintf("%s - TomNomNom.com", strings.Trim(fistLine, "# ")),
			Body:    template.HTML(blackfriday.MarkdownCommon(b)),
			DevMode: true,
		}
	}

	return posts
}

type PostsHandler struct {
	DevMode  bool
	Template *template.Template
	Posts    map[string]*Page
}

func (p *PostsHandler) GetPost(name string) (*Page, *PageError) {
	post, ok := p.Posts[name]
	if ok {
		return post, nil
	}

	// Check for old-style post names
	name = strings.Replace(name, "_", "-", -1)
	post, ok = p.Posts[name]
	if ok {
		return nil, &PageError{
			Message:  "Page moved",
			Code:     http.StatusMovedPermanently,
			Location: fmt.Sprintf("/posts/%s", name),
		}
	}

	return nil, &PageError{
		Message: "Page not found",
		Code:    http.StatusNotFound,
	}
}

func (p *PostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	name := parts[2]
	post, perr := p.GetPost(name)

	if perr != nil {
		switch perr.Code {
		case http.StatusMovedPermanently:
			http.Redirect(w, r, perr.Location, perr.Code)
		case http.StatusNotFound:
			http.Error(w, perr.Message, perr.Code)
		}
		return
	}

	err := p.Template.Execute(w, post)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}
}
