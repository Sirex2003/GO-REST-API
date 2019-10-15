package view

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorestapi/modules/datainit"
	"net/http"
	"strconv"
	"time"
)

//Subroutes list
func EventsRoutes(r *mux.Router) {
	r.HandleFunc("", Events).Methods(http.MethodGet)
	r.HandleFunc("/", Events).Methods(http.MethodGet)
	r.HandleFunc("/year/{year}", YearEvents).Methods(http.MethodGet)
	r.HandleFunc("/id/{id}", IDEvent).Methods(http.MethodGet)
}

// Events functions
func Events(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(datainit.EventsData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logrus.Println(err.Error())
		return
	}
}

//YearEvents
func YearEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	str, err := strconv.Atoi(params["year"])
	switch {
	case params["year"] == "":
		http.Error(w, "YearEvents can not be empty", http.StatusBadRequest)
		logrus.Println(r, "YearEvents set empty")
		return
	case err != nil:
		http.Error(w, "YearEvents must be a number", http.StatusBadRequest)
		logrus.Println(r, "YearEvents set as letters")
		return
	case 2000 < str || str > time.Now().Year():
		http.Error(w, "Invalid YearEvents value", http.StatusBadRequest)
		logrus.Println(r, "Invalid YearEvents value")
		return
	}
	var date time.Time
	for _, item := range datainit.EventsData {
		date, _ = time.Parse("2006-01-02", item.StartDate)
		if date.Format("2006") == params["year"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				logrus.Fatal(err.Error())
				return
			}
		}
	}
	http.Error(w, "YearEvents not found", http.StatusNoContent)
}

func IDEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		logrus.Println(err.Error())
		return
	}
	for _, item := range datainit.EventsData {
		if item.ID == params["id"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				logrus.Fatal(err.Error())
				return
			}
			return
		}
	}
	http.Error(w, "ID not found", http.StatusNoContent)
}
