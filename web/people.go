package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Person struct {
	Name      string     `json:"name"`
	Height    string     `json:"height"`
	Mass      string     `json:"mass"`
	HairColor string     `json:"hair_color"`
	SkinColor string     `json:"skin_color"`
	EyeColor  string     `json:"eye_color"`
	BirthYear string     `json:"birth_year"`
	Gender    string     `json:"gender"`
	Homeworld string     `json:"homeworld"`
	Films     []Film     `json:"films"`
	Species   []Species  `json:"species"`
	Vehicles  []Vehicle  `json:"vehicles"`
	Starships []Starship `json:"starships"`
	Metadata
}

type People struct {
	ResponseData
	Results []Person `json:"results"`
}

func initializePeopleRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/people").Subrouter()
	subrouter.HandleFunc("/", PeopleHandler)
	subrouter.HandleFunc("/{id}", PersonHandler)
}

func PeopleHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, People{
		ResponseData: ResponseData{},
		Results:      nil,
	})
}

func PersonHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, Person{
		Name:      "",
		Height:    "",
		Mass:      "",
		HairColor: "",
		SkinColor: "",
		EyeColor:  "",
		BirthYear: "",
		Gender:    "",
		Homeworld: "",
		Films:     nil,
		Species:   nil,
		Vehicles:  nil,
		Starships: nil,
		Metadata:  Metadata{},
	})
}
