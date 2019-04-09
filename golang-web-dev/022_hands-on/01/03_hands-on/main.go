package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func handleCat(w http.ResponseWriter, r *http.Request) {
	resp := "The Cat Says: Meow!"
	tpl.ExecuteTemplate(w, "index.gohtml", resp)
}

func handleDog(w http.ResponseWriter, r *http.Request) {
	resp := "The Dog Says: Woof!"
	tpl.ExecuteTemplate(w, "index.gohtml", resp)
}

func handleQuit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is now stopping!")
	defer func() { go log.Fatal("Shutting down server at request from ", r.RemoteAddr) }()
}

func main() {
	// This is really lazy and more difficult to read
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
		if err != nil {
			log.Fatal("Could not parse template: ", err)
		}
	})

	// Definitely more elegant
	http.HandleFunc("/cat", handleCat)
	http.HandleFunc("/dog", handleDog)
	http.HandleFunc("/quit", handleQuit)

	http.ListenAndServe(":8080", nil)
}
