package main

import (
	"github.com/gorilla/mux"
	"holonet/web"
	"log"
	"net/http"
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
	log.Fatal(http.ListenAndServe(appAddress, app.Router))
}
