package web

import (
	"database/sql"
	"github.com/gorilla/mux"
	"holonet/data/db"
	"holonet/data/resource"
	"log"
	"net/http"
	"strconv"
	"strings"
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
				Count:    0,
				Next:     "",
				Previous: "",
			},
			Results: films,
		}, 200)
	} else {
		writer.WriteHeader(http.StatusInternalServerError)

		log.Println(err)

		writeJSON(writer, Films{
			ResponseData: ResponseData{},
			Results:      films,
		}, 200)
	}
}

func FilmHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	filmId, err := strconv.Atoi(vars["id"])
	film, err := db.FindFilm(filmId)

	if err == nil {
		writeJSON(writer, film, 200)
	} else {
		var msg string
		var statusCode int
		if strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
			statusCode = http.StatusNotFound
			msg = "Not Found"
		} else {
			statusCode = http.StatusInternalServerError
			msg = "Server Error"
		}

		writeJSON(writer, ErrorResponse{Message: msg}, statusCode)
	}
}
