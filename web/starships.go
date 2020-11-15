package web

import (
	"github.com/gorilla/mux"
	"holonet/data/db"
	"holonet/data/resource"
	"net/http"
	"strconv"
)

type Starships struct {
	ResponseData
	Results []resource.Starship `json:"results"`
}

func initializeStarshipsRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/starships").Subrouter()
	subrouter.HandleFunc("/", StarshipsHandler)
	subrouter.HandleFunc("/{id}", StarshipHandler)
}

func StarshipsHandler(writer http.ResponseWriter, _ *http.Request) {
	starships, err := db.AllStarships()

	if err == nil {
		writeJSON(writer, Starships{
			ResponseData: ResponseData{
				Count:    len(starships),
				Next:     "",
				Previous: "",
			},
			Results: starships,
		}, 200)
	} else {
		handleIndexError(writer, err)
	}
}

func StarshipHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	starshipId, err := strconv.Atoi(vars["id"])
	starship, err := db.FindStarship(starshipId)

	if err == nil {
		writeJSON(writer, starship, 200)
	} else {
		handleFindError(writer, err)
	}
}
