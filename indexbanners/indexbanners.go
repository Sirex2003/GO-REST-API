package indexbanners

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type indexBanner struct {
	ID               string `json:"id"`
	EventName        string `json:"event_name"`
	EventDescription string `json:"event_description"`
	EventDates       string `json:"event_dates"`
	IsHidden         bool   `json:"is_visible"`
}

//TODO Подключить СУБД в качестве источника

//Init test data
var indexBanners []indexBanner

func init() {
	indexBanners = append(indexBanners, indexBanner{"1", "Highload 2019", "Highload++", "5-6 Ноября", false})
	indexBanners = append(indexBanners, indexBanner{"2", "Jocker 2019", "Jocker", "5-6 Сентябрь", true})
	indexBanners = append(indexBanners, indexBanner{"3", "РИТ 2019", "РИТ++", "5-6 Октябрь", false})
	indexBanners = append(indexBanners, indexBanner{"4", "МАКС 2019", "МАКС", "5-6 Декабря", true})
}

//Subroutes list
func Routes(subrouter *mux.Router) {
	subrouter.HandleFunc("", GetIndexBanners).Methods(http.MethodGet)
	subrouter.HandleFunc("/", GetIndexBanners).Methods(http.MethodGet)
	subrouter.HandleFunc("/visible", GetIndexBannersVisible).Methods(http.MethodGet)
	subrouter.HandleFunc("/{id}", UpdateIndexBanner).Methods(http.MethodGet, http.MethodPut, http.MethodPatch)
}

//TODO Resource manager to handle edit concurrency

//Index page banners
func GetIndexBanners(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(indexBanners); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logrus.Println(err.Error())
		return
	}
}

func GetIndexBannersVisible(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	i := 0
	for _, item := range indexBanners {
		if item.IsHidden == true {
			i++
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				logrus.Println(err.Error())
				return
			}
		}
	}
	if i == 0 {
		http.Error(w, "No visible banners found", http.StatusNoContent)
	}
}

func UpdateIndexBanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var banner indexBanner
	if _, err := strconv.Atoi(params["id"]); err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		logrus.Println(err.Error())
		return
	}
	switch r.Method {
	case http.MethodGet:
		for _, item := range indexBanners {
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
	case http.MethodPut:
		for index, item := range indexBanners {
			if item.ID == params["id"] {
				indexBanners = append(indexBanners[:index], indexBanners[index+1:]...)
				_ = json.NewDecoder(r.Body).Decode(&banner)
				banner.ID = params["id"]
				indexBanners = append(indexBanners, banner)
				if err := json.NewEncoder(w).Encode(banner); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					logrus.Println(err.Error())
					return
				}
				return
			}
		}
		http.Error(w, "ID not found", http.StatusBadRequest)
	case http.MethodPatch:
		_ = json.NewDecoder(r.Body).Decode(&banner)
		for index, item := range indexBanners {
			if item.ID == params["id"] {
				indexBanners[index].IsHidden = banner.IsHidden
				return
			}
		}
		http.Error(w, "ID not found", http.StatusBadRequest)
	}
}
