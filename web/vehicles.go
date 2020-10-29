package web

import (
	"github.com/gorilla/mux"
	"holonet/data"
	"net/http"
)

type Vehicles struct {
	data.Metadata
	Results []data.Vehicle `json:"results"`
}

func initializeVehiclesRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/vehicles").Subrouter()
	subrouter.HandleFunc("/", VehiclesHandler)
	subrouter.HandleFunc("/{id}", VehicleHandler)
}

func VehiclesHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, Vehicles{
		Metadata: data.Metadata{},
		Results:  nil,
	})
}

func VehicleHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, data.Vehicle{
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
		VehicleClass:         "",
		Pilots:               nil,
		Films:                nil,
		Metadata:             data.Metadata{},
	})
}
