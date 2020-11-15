package web

import (
	"github.com/gorilla/mux"
	"holonet/data/db"
	"holonet/data/resource"
	"net/http"
	"strconv"
)

type Vehicles struct {
	ResponseData
	Results []resource.Vehicle `json:"results"`
}

func initializeVehiclesRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/vehicles").Subrouter()
	subrouter.HandleFunc("/", VehiclesHandler)
	subrouter.HandleFunc("/{id}", VehicleHandler)
}

func VehiclesHandler(writer http.ResponseWriter, _ *http.Request) {
	vehicles, err := db.AllVehicles()

	if err == nil {
		writeJSON(writer, Vehicles{
			ResponseData: ResponseData{
				Count:    len(vehicles),
				Next:     "",
				Previous: "",
			},
			Results: vehicles,
		}, 200)
	} else {
		handleIndexError(writer, err)
	}
}

func VehicleHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	vehicleId, err := strconv.Atoi(vars["id"])
	vehicle, err := db.FindVehicle(vehicleId)

	if err == nil {
		writeJSON(writer, vehicle, 200)
	} else {
		handleFindError(writer, err)
	}
}
