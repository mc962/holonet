package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"holonet/web"
	"log"
	"net/http"
	"os"
)

const DefaultAddress = ":8010"

type Application struct {
	Router *mux.Router
}

func (app *Application) Initialize() {
	app.Router = mux.NewRouter().StrictSlash(true)

	web.InitializeRoutes(app.Router)
}

func (app *Application) Run(address string) {
	var appAddress string
	if len(address) > 0 {
		appAddress = address
	} else {
		appAddress = DefaultAddress
	}

	loggedRouter := handlers.LoggingHandler(os.Stdout, app.Router)

	log.Fatal(http.ListenAndServe(appAddress, loggedRouter))
}
