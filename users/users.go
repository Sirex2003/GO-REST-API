package users

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
)

type user struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

//TODO Подключить СУБД в качестве источника

//Init test data
var users []user

func init() {
	users = append(users, user{"1", "test", "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"})
}

//TODO Resource manager to handle edit concurrency

//Subroutes list
func Routes(subrouter *mux.Router) {
	subrouter.HandleFunc("", usersFunc).Methods(http.MethodGet, http.MethodPost)
	subrouter.HandleFunc("/", usersFunc).Methods(http.MethodGet, http.MethodPost)
	subrouter.HandleFunc("/{id}", updateUser).Methods(http.MethodPut)
}

func usersFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logrus.Println(err.Error())
			return
		}
	case http.MethodPost:
		var user user
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logrus.Println(err.Error())
			return
		}
		user.Id = strconv.Itoa(rand.Intn(100))
		users = append(users, user)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logrus.Println(err.Error())
			return
		}
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.Id == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user user
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.Id = params["id"]
			users = append(users, user)
			if err := json.NewEncoder(w).Encode(user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				logrus.Println(err.Error())
				return
			}
			return
		}
	}
	http.Error(w, "ID not found", http.StatusNoContent)
}
