package web

import (
	"github.com/gorilla/mux"
	"holonet/data/db"
	"holonet/data/resource"
	"net/http"
	"strconv"
)

type Planets struct {
	ResponseData
	Results []resource.Planet `json:"results"`
}

func initializePlanetsRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/planets").Subrouter()
	subrouter.HandleFunc("/", PlanetsHandler)
	subrouter.HandleFunc("/{id}", PlanetHandler)
}

func PlanetsHandler(writer http.ResponseWriter, _ *http.Request) {
	planets, err := db.AllPlanets()

	if err == nil {
		writeJSON(writer, Planets{
			ResponseData: ResponseData{
				Count:    len(planets),
				Next:     "",
				Previous: "",
			},
			Results: planets,
		}, 200)
	} else {
		handleIndexError(writer, err)
	}
}

func PlanetHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	planetId, err := strconv.Atoi(vars["id"])
	planet, err := db.FindPlanet(planetId)

	if err == nil {
		writeJSON(writer, planet, 200)
	} else {
		handleFindError(writer, err)
	}
}
