package web

import (
	"github.com/gorilla/mux"
	"holonet/data/resource"
	"net/http"
)

type Starships struct {
	ResponseData
	Results []resource.Starship `json:"results"`
}

func initializeStarshipsRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/starships").Subrouter()
	subrouter.HandleFunc("/", StarshipsHandler)
	subrouter.HandleFunc("/{id}", StarshipHandler)
}

func StarshipsHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, Starships{
		ResponseData: ResponseData{},
		Results:      nil,
	}, 200)
}

func StarshipHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, resource.Starship{
		Name:                 "",
		Model:                "",
		Manufacturer:         "",
		CostInCredits:        "",
		Length:               "",
		MaxAtmospheringSpeed: "",
		Crew:                 "",
		Passengers:           "",
		CargoCapacity:        "",
		Consumables:          "",
		HyperdriveRating:     "",
		MGLT:                 "",
		StarshipClass:        "",
		Pilots:               nil,
		Films:                nil,
		Metadata:             resource.Metadata{},
	}, 200)
}
