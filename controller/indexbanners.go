package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorestapi/modules/datainit"
	"net/http"
	"strconv"
)

//Subroutes list
func IndexbannersRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("/{id}", UpdateIndexBanner).Methods(http.MethodPut, http.MethodPatch)
}

//Index page banners
func UpdateIndexBanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var banner datainit.IndexBanner
	if _, err := strconv.Atoi(params["id"]); err != nil {
		http.Error(w, "ID is not a number", http.StatusBadRequest)
		logrus.Println(err.Error())
		return
	}
	switch r.Method {
	case http.MethodPut:
		for index, item := range datainit.IndexBannersData {
			if item.ID == params["id"] {
				datainit.IndexBannersData = append(datainit.IndexBannersData[:index], datainit.IndexBannersData[index+1:]...)
				_ = json.NewDecoder(r.Body).Decode(&banner)
				banner.ID = params["id"]
				datainit.IndexBannersData = append(datainit.IndexBannersData, banner)
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
		for index, item := range datainit.IndexBannersData {
			if item.ID == params["id"] {
				datainit.IndexBannersData[index].IsHidden = banner.IsHidden
				return
			}
		}
		http.Error(w, "ID not found", http.StatusBadRequest)
	}
}
