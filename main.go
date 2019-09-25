package main

import (
	"./authentication"
	"./contacts"
	"./events"
	"./indexbanners"
	"./users"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//TODO Испробовать logrus в качестве логгера

func main() {
	router := mux.NewRouter()
	//Auth middleware
	router.Use(authentication.Authentication)

	//Events Subrouter
	events.Routes(router.PathPrefix("/events").Subrouter())
	//Index banners Subrouter
	indexbanners.Routes(router.PathPrefix("/indexbanners").Subrouter())
	//Contacts Subrouter
	contacts.Routes(router.PathPrefix("/contacts").Subrouter())
	//Users Subrouter
	users.Routes(router.PathPrefix("/users").Subrouter())

	// TODO Подключить TLS
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
