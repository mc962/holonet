package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Film struct {
	Title        string     `json:"title"`
	EpisodeId    int        `json:"episode_id"`
	OpeningCrawl string     `json:"opening_crawl"`
	Director     string     `json:"director"`
	Producer     string     `json:"producer"`
	ReleaseDate  string     `json:"release_date"`
	Characters   []Person   `json:"characters"`
	Planets      []Planet   `json:"planets"`
	Starships    []Starship `json:"starships"`
	Vehicles     []Vehicle  `json:"vehicles"`
	Species      []Species  `json:"species"`
	Metadata
}

func initializeFilmsRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/films").Subrouter()
	subrouter.HandleFunc("/", FilmsHandler)
	subrouter.HandleFunc("/{id}", FilmHandler)
}

type Films struct {
	ResponseData
	Results []Film `json:"results"`
}

func FilmsHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, Films{
		ResponseData: ResponseData{},
		Results:      nil,
	})
}

func FilmHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, Film{
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
		Metadata:     Metadata{},
	})
}
