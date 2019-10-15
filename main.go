package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorestapi/controller"
	"gorestapi/modules/datainit"
	"gorestapi/modules/middleware"
	"gorestapi/view"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//TODO Migrate to standart "net/http"
	datainit.InitTestData()
	//Router
	router := mux.NewRouter()
	//Middleware
	router.Use(middleware.Authentication)
	router.Use(middleware.RequestLog)

	//Sub-Routers @Controllers
	controller.EventsRoutes(router.Methods(http.MethodPost, http.MethodPut, http.MethodDelete).PathPrefix("/events").Subrouter())
	controller.IndexbannersRoutes(router.Methods(http.MethodPut, http.MethodPatch).PathPrefix("/indexbanners").Subrouter())
	controller.ContactsRoutes(router.Methods(http.MethodPut).PathPrefix("/contacts").Subrouter())
	controller.UsersRoutes(router.Methods(http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete).PathPrefix("/users").Subrouter())

	//Sub-Routers @Views
	view.EventsRoutes(router.Methods(http.MethodGet).PathPrefix("/events").Subrouter())
	view.IndexBannersRoutes(router.Methods(http.MethodGet).PathPrefix("/indexbanners").Subrouter())
	view.ContactsRoutes(router.Methods(http.MethodGet).PathPrefix("/contacts").Subrouter())
	view.UsersRoutes(router.Methods(http.MethodGet).PathPrefix("/users").Subrouter())

	//Get Heroku port for Web <BEGIN>
	//port := os.Getenv("PORT")

	port := "8000"
	//if port == "" {
	//	logrus.Fatal("$PORT must be set")
	//}
	//Get Heroku port for Web <END>

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
		if err := srv.ListenAndServeTLS("./modules/certs/server.crt", "./modules/certs/server.key"); err != nil {
			logrus.WithFields(logrus.Fields{"port": port}).Error(err.Error())
		}
	}()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{"port": port}).Info("Server is starting up")

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
