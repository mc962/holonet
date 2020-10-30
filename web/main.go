package web

import "github.com/gorilla/mux"

func Initialize() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router = RegisterRoutes(router)

	return router
}

func RegisterRoutes(router *mux.Router) *mux.Router {
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

	return router
}
