package view

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorestapi/modules/datainit"
	"net/http"
	"strconv"
)

//Subroutes list
func IndexBannersRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("", GetIndexBanners).Methods(http.MethodGet)
	subrouter.HandleFunc("/", GetIndexBanners).Methods(http.MethodGet)
	subrouter.HandleFunc("/visible", GetIndexBannersVisible).Methods(http.MethodGet)
	subrouter.HandleFunc("/{id}", GetIndexBannerId).Methods(http.MethodGet)
}

//TODO Resource manager to handle edit concurrency

//Index page banners
func GetIndexBanners(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(datainit.IndexBannersData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logrus.Println(err.Error())
		return
	}
}

func GetIndexBannersVisible(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	i := 0
	for _, item := range datainit.IndexBannersData {
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

func GetIndexBannerId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		logrus.Println(err.Error())
		return
	}
	for _, item := range datainit.IndexBannersData {
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
