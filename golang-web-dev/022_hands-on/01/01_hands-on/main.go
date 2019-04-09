package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleCat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Meow!")

}

func handleDog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Woof!")
}

func handleQuit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is now stopping!")
	defer func() { go log.Fatal("Shutting down server at request from ", r.RemoteAddr) }()
}

func main() {
	// This is really lazy and more difficult to read
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Stuff!")
	})

	// Definitely more elegant
	http.HandleFunc("/cat", handleCat)
	http.HandleFunc("/dog", handleDog)
	http.HandleFunc("/quit", handleQuit)

	http.ListenAndServe(":8080", nil)
}
