package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []Person `json:"residents"`
	Films          []Film   `json:"films"`
	Metadata
}

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
	})
}

func PlanetHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, Planet{
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
		Metadata:       Metadata{},
	})
}
