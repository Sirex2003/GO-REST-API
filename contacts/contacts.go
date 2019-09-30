package contacts

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type contacts struct {
	Phone1  string `json:"phone1"`
	Phone2  string `json:"phone2"`
	Address string `json:"address"`
	WebURL  string `json:"web_url"`
	Email   string `json:"email"`
	MapCode string `json:"map_code"`
}

//TODO Подключить СУБД в качестве источника

//Init test data
var contact contacts

func init() {
	contact.Phone1 = "109857438946"
	contact.Phone2 = "109857438946"
	contact.Address = "Москва, Фиг-знает где 48/90"
	contact.Email = "mail@example.com"
	contact.WebURL = "https://example.com"
	contact.MapCode = "https://example.com/map?=12kflankjshiu34rcoqiy"
}

//Subroutes list
func Routes(subrouter *mux.Router) {
	subrouter.HandleFunc("", Contacts).Methods(http.MethodGet, http.MethodPut)
	subrouter.HandleFunc("/", Contacts).Methods(http.MethodGet, http.MethodPut)
}

//TODO Resource manager to handle edit concurrency

func Contacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(contact); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logrus.Println(err.Error())
			return
		}
	case http.MethodPut:
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logrus.Println(err.Error())
			return
		}
	}
}
