package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Starship struct {
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
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	MGLT                 string   `json:"mglt"`
	StarshipClass        string   `json:"starship_class"`
	Pilots               []Person `json:"pilots"`
	Films                []Film   `json:"films"`
	Metadata
}

type Starships struct {
	ResponseData
	Results []Starship `json:"results"`
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
	writeJSON(writer, Starship{
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
		Metadata:             Metadata{},
	})
}
