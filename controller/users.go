package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorestapi/modules/datainit"
	"math/rand"
	"net/http"
	"strconv"
)

//Subroutes list
func UsersRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("", usersFunc).Methods(http.MethodPost)
	subrouter.HandleFunc("/", usersFunc).Methods(http.MethodPost)
	subrouter.HandleFunc("/{id}", updateUser).Methods(http.MethodPut)
}

//Func users
func usersFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user datainit.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logrus.Println(err.Error())
		return
	}
	user.Id = strconv.Itoa(rand.Intn(100))
	datainit.UsersData = append(datainit.UsersData, user)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logrus.Println(err.Error())
		return
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range datainit.UsersData {
		if item.Id == params["id"] {
			datainit.UsersData = append(datainit.UsersData[:index], datainit.UsersData[index+1:]...)
			var user datainit.User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.Id = params["id"]
			datainit.UsersData = append(datainit.UsersData, user)
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
