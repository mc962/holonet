package web

import (
	"github.com/gorilla/mux"
	"holonet/data/resource"
	"net/http"
)

type Planets struct {
	ResponseData
	Results []Planets `json:"results"`
}

func initializePlanetsRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/planets").Subrouter()
	subrouter.HandleFunc("/", PlanetsHandler)
	subrouter.HandleFunc("/{id}", PlanetHandler)
}

func PlanetsHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, Planets{
		ResponseData: ResponseData{},
		Results:      nil,
	}, 200)
}

func PlanetHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, resource.Planet{
		Name:           "",
		RotationPeriod: "",
		OrbitalPeriod:  "",
		Diameter:       "",
		Climate:        "",
		Gravity:        "",
		Terrain:        "",
		SurfaceWater:   "",
		Population:     "",
		Residents:      nil,
		Films:          nil,
		Metadata:       resource.Metadata{},
	}, 200)
}
