package web

import (
	"github.com/gorilla/mux"
	"holonet/data"
	"net/http"
)

type ManySpecies struct {
	ResponseData
	Results []data.Species `json:"results"`
}

func initializeSpeciesRouter(router *mux.Router) {
	subrouter := router.PathPrefix("/species").Subrouter()
	subrouter.HandleFunc("/", AllSpeciesHandler)
	subrouter.HandleFunc("/{id}", SpeciesHandler)
}

func AllSpeciesHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, ManySpecies{
		ResponseData: ResponseData{},
		Results:      nil,
	})
}

func SpeciesHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, data.Species{
		Name:            "",
		Classification:  "",
		Designation:     "",
		AverageHeight:   "",
		AverageLifespan: "",
		HairColors:      "",
		SkinColors:      "",
		EyeColors:       "",
		Homeworld:       "",
		Language:        "",
		People:          nil,
		Films:           nil,
		Metadata:        data.Metadata{},
	})
}
