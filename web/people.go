package web

import (
	"github.com/gorilla/mux"
	"holonet/data/db"
	"holonet/data/resource"
	"net/http"
	"strconv"
)

type People struct {
	ResponseData
	Results []resource.Person `json:"results"`
}

func initializePeopleRoutes(router *mux.Router) {
	subrouter := router.PathPrefix("/people").Subrouter()
	subrouter.HandleFunc("/", PeopleHandler)
	subrouter.HandleFunc("/{id}", PersonHandler)
}

func PeopleHandler(writer http.ResponseWriter, _ *http.Request) {
	people, err := db.AllPeople()

	if err == nil {
		writeJSON(writer, People{
			ResponseData: ResponseData{
				Count:    len(people),
				Next:     "",
				Previous: "",
			},
			Results: people,
		}, 200)
	} else {
		handleIndexError(writer, err)
	}
}

func PersonHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	personId, err := strconv.Atoi(vars["id"])
	person, err := db.FindPerson(personId)

	if err == nil {
		writeJSON(writer, person, 200)
	} else {
		handleFindError(writer, err)
	}
}
