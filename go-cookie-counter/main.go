package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("My-Counter")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "My-Counter",
			Value: "0",
		}
	}

	counter, _ := strconv.Atoi(c.Value)
	counter++
	c.Value = strconv.Itoa(counter)

	// Ensure Cookie is set before we send the body
	http.SetCookie(w, c)

	// Send the body
	fmt.Fprintln(w, "<h1>Welcome!</h1>")
	if counter > 1 {
		fmt.Fprintln(w, "You've been here ", c.Value, " times before")
		fmt.Printf("%T\n", c)
		fmt.Printf("%v\n", c)
	} else {
		fmt.Fprintln(w, "This is your first time here!")
	}

}

func main() {
	http.HandleFunc("/", handleRoot)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
