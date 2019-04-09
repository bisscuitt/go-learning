package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func dogImg(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpeg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog.jpeg", dogImg)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}
