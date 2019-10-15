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
func EventsRoutes(r *mux.Router) {
	r.HandleFunc("", Events).Methods(http.MethodPost)
	r.HandleFunc("/", Events).Methods(http.MethodPost)
	r.HandleFunc("/id/{id}", IDEvent).Methods(http.MethodPut, http.MethodDelete)
}

// Events functions
func Events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(Events); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logrus.Println(err.Error())
			return
		}
	case http.MethodPost:
		var event datainit.Event
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logrus.Println(err.Error())
			return
		}
		event.ID = strconv.Itoa(rand.Intn(100))
		datainit.EventsData = append(datainit.EventsData, event)
		if err := json.NewEncoder(w).Encode(event); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logrus.Println(err.Error())
			return
		}
	}
}

func IDEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		logrus.Println(err.Error())
		return
	}
	switch r.Method {
	case http.MethodPut:
		for index, item := range datainit.EventsData {
			if item.ID == params["id"] {
				datainit.EventsData = append(datainit.EventsData[:index], datainit.EventsData[index+1:]...)
				var event datainit.Event
				_ = json.NewDecoder(r.Body).Decode(&event)
				event.ID = params["id"]
				datainit.EventsData = append(datainit.EventsData, event)
				if err := json.NewEncoder(w).Encode(event); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					logrus.Println(err.Error())
					return
				}
				return
			}
		}
		http.Error(w, "ID not found", http.StatusNoContent)
	case http.MethodDelete:
		for index, item := range datainit.EventsData {
			if item.ID == params["id"] {
				datainit.EventsData = append(datainit.EventsData[:index], datainit.EventsData[index+1:]...)
				break
			}
		}
		http.Error(w, "ID not found", http.StatusNoContent)
	}
}
