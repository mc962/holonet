package web

import (
	"github.com/gorilla/mux"
	"holonet/data"
	"net/http"
)

type People struct {
	ResponseData
	Results []data.Person `json:"results"`
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
	writeJSON(writer, data.Person{
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
		Metadata:  data.Metadata{},
	})
}
