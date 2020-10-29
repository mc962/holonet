package web

import (
	"github.com/gorilla/mux"
	"holonet/data"
	"net/http"
)

type Starships struct {
	ResponseData
	Results []data.Starship `json:"results"`
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
	})
}

func StarshipHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, data.Starship{
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
		Metadata:             data.Metadata{},
	})
}
