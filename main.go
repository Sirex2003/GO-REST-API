package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorestapi/contacts"
	"gorestapi/events"
	"gorestapi/indexbanners"
	"gorestapi/middleware"
	"gorestapi/users"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	//Auth middleware
	router.Use(middleware.Authentication)
	router.Use(middleware.RequestLog)

	//Events Subrouter
	events.Routes(router.PathPrefix("/events").Subrouter())
	//Index banners Subrouter
	indexbanners.Routes(router.PathPrefix("/indexbanners").Subrouter())
	//Contacts Subrouter
	contacts.Routes(router.PathPrefix("/contacts").Subrouter())
	//Users Subrouter
	users.Routes(router.PathPrefix("/users").Subrouter())

	//Get Heroku port for Web <BEGIN>
	//port := os.Getenv("PORT")
	logrus.SetFormatter(&logrus.JSONFormatter{})
	port := "8000"
	if port == "" {
		logrus.Fatal("$PORT must be set")
	}
	//Get Heroku port for Web <END>
	logrus.WithFields(logrus.Fields{"port": port}).Info("Starting server")
	// TODO Подключить TLS
	logrus.Fatal(http.ListenAndServe(":"+port, router))
}
