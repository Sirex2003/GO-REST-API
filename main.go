package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorestapi/contacts"
	"gorestapi/events"
	"gorestapi/indexbanners"
	"gorestapi/middleware"
	"gorestapi/users"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
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
	//if port == "" {
	//	logrus.Fatal("$PORT must be set")
	//}
	//Get Heroku port for Web <END>
	logrus.WithFields(logrus.Fields{"port": port}).Info("Server is starting up")

	//HTTP-Server with timeouts
	srv := &http.Server{
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServeTLS("./certs/server.crt", "./certs/server.key"); err != nil {
			logrus.WithFields(logrus.Fields{"port": port}).Error(err.Error())
		}
	}()

	//Channel with buffer size = 1, for waiting shutdown command
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// A deadline to wait for.
	wait := time.Second * 120
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	logrus.WithFields(logrus.Fields{"port": port}).Info("Server is shutting down")
	os.Exit(0)
}
