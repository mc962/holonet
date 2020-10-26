package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeRoutes(router *mux.Router) {
	// top-level routes
	router.HandleFunc("/", RootHandler)
	subrouter := router.PathPrefix("/api").Subrouter().StrictSlash(true)
	subrouter.HandleFunc("/", APIRootHandler)

	// namespaced resource routes
	initializeFilmsRoutes(subrouter)
	initializePeopleRoutes(subrouter)
	initializePlanetsRoutes(subrouter)
	initializeSpeciesRouter(subrouter)
	initializeStarshipsRoutes(subrouter)
	initializeVehiclesRoutes(subrouter)
}

func RootHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Accept", "application/json")
	_ = json.NewEncoder(writer).Encode(map[string]string{
		"version": "0.1.0",
	})
}

func APIRootHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Accept", "application/json")

	siteRoutes := map[string]string{
		"people":    "http://holonet.ai/api/people/",
		"planets":   "http://holonet.ai/api/planets/",
		"films":     "http://holonet.ai/api/films/",
		"species":   "http://holonet.ai/api/species/",
		"vehicles":  "http://holonet.ai/api/vehicles/",
		"starships": "http://holonet.ai/api/starships/",
	}

	writeJSON(writer, siteRoutes)
}
