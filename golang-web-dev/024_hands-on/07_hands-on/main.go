package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./starting-files/templates/*.gohtml"))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	fs := http.FileServer(http.Dir("./starting-files/public"))

	http.HandleFunc("/", handleRoot)
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	http.ListenAndServe(":8080", nil)
}
