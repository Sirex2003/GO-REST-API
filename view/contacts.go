package view

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorestapi/modules/datainit"
	"net/http"
)

//Subroutes list
func ContactsRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("", Contacts).Methods(http.MethodGet)
	subrouter.HandleFunc("/", Contacts).Methods(http.MethodGet)
}

//Func contacts
func Contacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(datainit.ContactData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logrus.Println(err.Error())
		return
	}
}
