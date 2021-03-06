package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	url2 "net/url"
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

func APIRootHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Accept", "application/json")

	siteRoute := func(path string) string {
		tlsStatus := request.TLS
		var urlScheme string

		if tlsStatus == nil {
			urlScheme = "http"
		} else {
			urlScheme = "https"
		}

		var routeURL string

		url, err := url2.Parse(path)
		if err == nil {
			url.Scheme = urlScheme
			url.Host = request.Host

			routeURL = url.String()
		} else {
			routeURL = "#"
		}

		return routeURL
	}

	siteRoutes := map[string]string{
		"people":    siteRoute("/api/people/"),
		"planets":   siteRoute("/api/planets/"),
		"films":     siteRoute("/api/films/"),
		"species":   siteRoute("/api/species/"),
		"vehicles":  siteRoute("/api/vehicles/"),
		"starships": siteRoute("/api/starships/"),
	}

	writeJSON(writer, siteRoutes)
}
