package main

import (
	"./contacts"
	"./events"
	"./indexbanners"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

/*
type user struct {
	Login        string `json:"login"`
	PasswordHash string `json:"hash"`
}
*/

func main() {
	router := mux.NewRouter()
	//init test data
	events.DataInit()
	indexbanners.DataInit()
	contacts.DataInit()

	//Events handlers
	router.HandleFunc("/events", events.GetEvents).Methods("GET")
	router.HandleFunc("/events/year/{year}", events.GetEventsYear).Methods("GET")
	router.HandleFunc("/events/id/{id}", events.GetEvent).Methods("GET")
	router.HandleFunc("/events", events.CreateEvent).Methods("POST")
	router.HandleFunc("/events/id/{id}", events.UpdateEvent).Methods("PUT")
	router.HandleFunc("/events/id/{id}", events.DeleteEvent).Methods("DELETE")

	//Index banners handlers
	router.HandleFunc("/indexbanners", indexbanners.GetIndexBanners).Methods("GET")
	router.HandleFunc("/indexbanners/visible", indexbanners.GetIndexBannersVisible).Methods("GET")
	router.HandleFunc("/indexbanners/{id}", indexbanners.UpdateIndexBanner).Methods("PUT")
	router.HandleFunc("/indexbanners/{id}", indexbanners.UpdateIndexBannersVisibility).Methods("PATCH")

	//Contacts handlers (Функционал для контактов)
	router.HandleFunc("/contacts", contacts.GetContacts).Methods("GET")
	router.HandleFunc("/contacts", contacts.UpdateContacts).Methods("PUT")

	/*
		//Users handlers (Функционал для пользователей)
		router.HandleFunc("/users", getUsers).Methods("GET")
		router.HandleFunc("/users", createUser).Methods("POST")
		router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	*/
	log.Fatal(http.ListenAndServe(":8000", router))
}
