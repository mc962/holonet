package web

import (
	"github.com/gorilla/mux"
	"holonet/data/resource"
	"net/http"
)

func initializeFilmsRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/films").Subrouter()
	subrouter.HandleFunc("/", FilmsHandler)
	subrouter.HandleFunc("/{id}", FilmHandler)
}

type Films struct {
	ResponseData
	Results []resource.Film `json:"results"`
}

func FilmsHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, Films{
		ResponseData: ResponseData{},
		Results:      nil,
	})
}

func FilmHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, resource.Film{
		Title:        "",
		EpisodeId:    0,
		OpeningCrawl: "",
		Director:     "",
		Producer:     "",
		ReleaseDate:  "",
		Characters:   nil,
		Planets:      nil,
		Starships:    nil,
		Vehicles:     nil,
		Species:      nil,
		Metadata:     resource.Metadata{},
	})
}
