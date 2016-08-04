package main

import (
	"net/http"
	"text/template"
	"sync"
	"path/filepath"
	"fmt"
)

type templateHandler struct {
	once sync.Once
	templ *template.Template
	filename string
}

func(t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func(){
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, nil)
}

func main() {
	room := &room{
		forward: make(chan []byte),
		join: make(chan *client),
		leave: make(chan *client),
		clients: make(map[*client]bool),
	}
	http.Handle("/room", room)
	go room.run()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	fmt.Println("Serving at port 8080")
	http.ListenAndServe(":8080", nil);
}
