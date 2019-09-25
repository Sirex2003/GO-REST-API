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

	//Events Subrouter
	events.Routes(router.PathPrefix("/events").Subrouter())
	//Index banners Subrouter
	indexbanners.Routes(router.PathPrefix("/indexbanners").Subrouter())
	//Contacts Subrouter
	contacts.Routes(router.PathPrefix("/contacts").Subrouter())
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
