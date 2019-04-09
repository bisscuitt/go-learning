package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bisscuitt/go-learning/golang-web-dev/042_mongodb/08_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(m map[string]models.User) *UserController {
	loadUsersFromFile(&m)

	return &UserController{m}

}

func loadUsersFromFile(m *map[string]models.User) {
	// Open the userfile (create iif not existing)
	f, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0640)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	jsonUsers, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// Move on if file is empty
	if len(jsonUsers) <= 0 {
		return
	}

	err = json.Unmarshal(jsonUsers, m)
	if err != nil {
		log.Fatal(err)
	}
}

func saveUsersToFile(m *map[string]models.User) {

	// Open the userfile (create if not existing)
	f, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0640)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	json_out, _ := json.Marshal(*m)
	f.Write(json_out)
}

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uj, _ := json.Marshal(uc.session)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Fetch user
	if _, ok := uc.session[id]; !ok {
		w.WriteHeader(404)
		return
	}

	u := uc.session[id]

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create new UUID
	u.Id = uuid.Must(uuid.NewV4()).String()

	// store the user in mongodb
	uc.session[u.Id] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)

	saveUsersToFile(&uc.session)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Delete user
	if _, ok := uc.session[id]; ok {
		w.WriteHeader(404)
		return
	}

	delete(uc.session, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")

	saveUsersToFile(&uc.session)
}
