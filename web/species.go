package web

import (
	"github.com/gorilla/mux"
	"holonet/data/db"
	"holonet/data/resource"
	"net/http"
	"strconv"
)

type ManySpecies struct {
	ResponseData
	Results []resource.Species `json:"results"`
}

func initializeSpeciesRouter(router *mux.Router) {
	subrouter := router.PathPrefix("/species").Subrouter()
	subrouter.HandleFunc("/", AllSpeciesHandler)
	subrouter.HandleFunc("/{id}", SpeciesHandler)
}

func AllSpeciesHandler(writer http.ResponseWriter, _ *http.Request) {
	species, err := db.AllSpecies()

	if err == nil {
		writeJSON(writer, ManySpecies{
			ResponseData: ResponseData{
				Count:    len(species),
				Next:     "",
				Previous: "",
			},
			Results: species,
		}, 200)
	} else {
		handleIndexError(writer, err)
	}
}

func SpeciesHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	speciesId, err := strconv.Atoi(vars["id"])
	species, err := db.FindSpecies(speciesId)

	if err == nil {
		writeJSON(writer, species, 200)
	} else {
		handleFindError(writer, err)
	}
}
