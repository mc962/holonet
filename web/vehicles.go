package web

import (
	"github.com/gorilla/mux"
	"holonet/data/resource"
	"net/http"
)

type Vehicles struct {
	resource.Metadata
	Results []resource.Vehicle `json:"results"`
}

func initializeVehiclesRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/vehicles").Subrouter()
	subrouter.HandleFunc("/", VehiclesHandler)
	subrouter.HandleFunc("/{id}", VehicleHandler)
}

func VehiclesHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, Vehicles{
		Metadata: resource.Metadata{},
		Results:  nil,
	}, 200)
}

func VehicleHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, resource.Vehicle{
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
		Metadata:             resource.Metadata{},
	}, 200)
}
