package view

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorestapi/modules/datainit"
	"net/http"
)

//Subroutes list
func UsersRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("", getUsers).Methods(http.MethodGet)
	subrouter.HandleFunc("/", getUsers).Methods(http.MethodGet)
	subrouter.HandleFunc("/{id}", getUserId).Methods(http.MethodGet)
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(datainit.UsersData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logrus.Println(err.Error())
		return
	}
}

func getUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range datainit.UsersData {
		if item.Id == params["id"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				logrus.Println(err.Error())
				return
			}
			return
		}
	}
	http.Error(w, "User not found", http.StatusNoContent)
}
