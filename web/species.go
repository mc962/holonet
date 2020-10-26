package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Species struct {
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	AverageLifespan string   `json:"average_lifespan"`
	HairColors      string   `json:"hair_colors"`
	SkinColors      string   `json:"skin_colors"`
	EyeColors       string   `json:"eye_colors"`
	Homeworld       string   `json:"homeworld"`
	Language        string   `json:"language"`
	People          []Person `json:"people"`
	Films           []Film   `json:"films"`
	Metadata
}

type ManySpecies struct {
	ResponseData
	Results []Species `json:"results"`
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
	writeJSON(writer, Species{
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
		Metadata:        Metadata{},
	})
}
