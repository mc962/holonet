package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/api", APIRootHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func RootHandler(writer http.ResponseWriter, _request *http.Request) {
	http.Redirect(writer, _request, "/api", http.StatusMovedPermanently)
}

func APIRootHandler(writer http.ResponseWriter, _request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(map[string]string{
		"version": "0.1.0",
	})
}