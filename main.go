package main

import (
	"./authentication"
	"./contacts"
	"./events"
	"./indexbanners"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	//init test data
	events.DataInit()
	indexbanners.DataInit()
	contacts.DataInit()
	authentication.DataInit()
	//TODO Испробовать logrus в качестве логгера

	router.Use(authentication.Authentication)

	//Events handlers
	router.HandleFunc("/events", events.Events).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/events/year/{year}", events.YearEvents).Methods(http.MethodGet)
	router.HandleFunc("/events/id/{id}", events.IDEvent).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	//Index banners handlers
	router.HandleFunc("/indexbanners", indexbanners.GetIndexBanners).Methods(http.MethodGet)
	router.HandleFunc("/indexbanners/visible", indexbanners.GetIndexBannersVisible).Methods(http.MethodGet)
	router.HandleFunc("/indexbanners/{id}", indexbanners.UpdateIndexBanner).Methods(http.MethodGet, http.MethodPut, http.MethodPatch)

	//Contacts handlers (Функционал для контактов)
	router.HandleFunc("/contacts", contacts.Contacts).Methods(http.MethodGet, http.MethodPut)
	/*
		TODO
			//Users handlers (Функционал для пользователей)
			router.HandleFunc("/users", getUsers).Methods("GET")
			router.HandleFunc("/users", createUser).Methods("POST")
			router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	*/

	// TODO Подключить TLS
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
