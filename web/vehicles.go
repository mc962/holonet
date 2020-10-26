package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Vehicle struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	VehicleClass         string   `json:"vehicle_class"`
	Pilots               []Person `json:"pilots"`
	Films                []Film   `json:"films"`
	Metadata
}

type Vehicles struct {
	Metadata
	Results []Vehicle `json:"results"`
}

func initializeVehiclesRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/vehicles").Subrouter()
	subrouter.HandleFunc("/", VehiclesHandler)
	subrouter.HandleFunc("/{id}", VehicleHandler)
}

func VehiclesHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, Vehicles{
		Metadata: Metadata{},
		Results:  nil,
	})
}

func VehicleHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, Vehicle{
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
		Metadata:             Metadata{},
	})
}
