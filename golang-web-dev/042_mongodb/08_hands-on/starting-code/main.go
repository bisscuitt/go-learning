package main

import (
	"github.com/bisscuitt/go-learning/golang-web-dev/042_mongodb/08_hands-on/starting-code/controllers"
	"github.com/bisscuitt/go-learning/golang-web-dev/042_mongodb/08_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()

	userMap := make(map[string]models.User)

	// Get a UserController instance
	uc := controllers.NewUserController(userMap)
	r.GET("/user/:id", uc.GetUser)
	r.GET("/users", uc.GetUsers)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
