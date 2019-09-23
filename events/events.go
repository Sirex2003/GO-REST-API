package events

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type event struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Header      string `json:"header"`
	Description string `json:"description"`
	Address     string `json:"address"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	ExternalURL string `json:"external_url"`
}

var events []event

func DataInit() {
	events = append(events, event{ID: "1", Name: "NTMEX 2003", Header: "NANO TECHNOLOGY EXPO", Description: "Нанотехнологии", Address: "Moscow", StartDate: "2003-12-20", EndDate: "2003-12-25", ExternalURL: "http://ntmex.ru"})
	events = append(events, event{ID: "2", Name: "FESTIVAL NAUKI 2003", Header: "FESTIVAL NAUKI", Description: "Фестиваль науки", Address: "Moscow", StartDate: "2003-11-20", EndDate: "2003-11-25", ExternalURL: "http://festivalnauki.ru"})
	events = append(events, event{ID: "3", Name: "МАКС 2003", Header: "МАКС", Description: "Международный авиасалон", Address: "Moscow", StartDate: "2003-10-20", EndDate: "2003-10-25", ExternalURL: "http://maks.ru"})
	events = append(events, event{ID: "4", Name: "HIGHLOAD 2003", Header: "HIGHLOAD", Description: "Высоконагруженные ИТ-системы", Address: "Moscow", StartDate: "2003-09-20", EndDate: "2003-09-25", ExternalURL: "http://highload.ru"})
}

// Events functions
func GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
}

func GetEventsYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	str, err := strconv.Atoi(params["year"])
	switch {
	case params["year"] == "":
		http.Error(w, "Year can not be empty", http.StatusBadRequest)
		log.Println(r, "Year set empty")
		return
	case err != nil:
		http.Error(w, "Year must be a number", http.StatusBadRequest)
		log.Println(r, "Year set as letters")
		return
	case 2000 < str || str > time.Now().Year():
		http.Error(w, "Invalid Year value", http.StatusBadRequest)
		log.Println(r, "Invalid Year value")
		return
	}

	if params["year"] == "" {

	}
	var date time.Time
	for _, item := range events {
		date, _ = time.Parse("2006-01-02", item.StartDate)
		if date.Format("2006") == params["year"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Fatal(err.Error())
				return
			}
		}
	}
	http.Error(w, "Year not found", http.StatusNoContent)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	for _, item := range events {
		if item.ID == params["id"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Fatal(err.Error())
				return
			}
			return
		}
	}
	http.Error(w, "ID not found", 204)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	event.ID = strconv.Itoa(rand.Intn(100))
	events = append(events, event)
	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	for index, item := range events {
		if item.ID == params["id"] {
			events = append(events[:index], events[index+1:]...)
			var event event
			_ = json.NewDecoder(r.Body).Decode(&event)
			event.ID = params["id"]
			events = append(events, event)
			if err := json.NewEncoder(w).Encode(event); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err.Error())
				return
			}
			return
		}
	}
	http.Error(w, "ID not found", http.StatusNoContent)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	for index, item := range events {
		if item.ID == params["id"] {
			events = append(events[:index], events[index+1:]...)
			break
		}
	}
}
