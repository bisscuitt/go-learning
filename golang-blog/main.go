package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

type output struct {
	Sid    string
	Images []string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/set/images", handleSetImage)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	// Check cookie for Session-ID
	sID := getSessionID(w, req)
	iCookie, _ := req.Cookie("Images")

	images := strings.Split(iCookie.Value, "|")

	data := output{
		Sid:    sID,
		Images: images,
	}

	log.Println(tpl.ExecuteTemplate(w, "index.gohtml", data))
}

func handleSetImage(w http.ResponseWriter, req *http.Request) {
	_ = getSessionID(w, req)

	ic, err := req.Cookie("Images")

	if err != nil {
		ic = &http.Cookie{
			Name:     "Images",
			Value:    "dog.jpg|test.jpg|woot.gif",
			HttpOnly: true,
			Path:     "/",
		}

	}
	http.SetCookie(w, ic)
}

func getSessionID(w http.ResponseWriter, req *http.Request) string {
	session_cookie, err := req.Cookie("Session-ID")

	if err != nil {
		uuid, _ := uuid.NewV4()
		// Generate Session-ID Cookie if not exist
		session_cookie = &http.Cookie{
			Name:     "Session-ID",
			Value:    uuid.String(),
			HttpOnly: true,
			Path:     "/",
		}
	}

	// Re-Set Expiry on Cookie
	session_cookie.MaxAge = 3600
	http.SetCookie(w, session_cookie)

	return session_cookie.Value
}
