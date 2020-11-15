package web

import (
	"github.com/gorilla/mux"
	"holonet/data/db"
	"holonet/data/resource"
	"net/http"
	"strconv"
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
	films, err := db.AllFilms()

	if err == nil {
		writeJSON(writer, Films{
			ResponseData: ResponseData{
				Count:    len(films),
				Next:     "",
				Previous: "",
			},
			Results: films,
		}, 200)
	} else {
		handleIndexError(writer, err)
	}
}

func FilmHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	filmId, err := strconv.Atoi(vars["id"])
	film, err := db.FindFilm(filmId)

	if err == nil {
		writeJSON(writer, film, 200)
	} else {
		handleFindError(writer, err)
	}
}
