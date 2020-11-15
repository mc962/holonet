package web

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
	"strings"
)

func writeJSON(writer http.ResponseWriter, response interface{}, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Accept", "application/json")
	writer.WriteHeader(statusCode)

	_ = json.NewEncoder(writer).Encode(response)
}

func navPages(request *http.Request, path string, currentPage string) (string string) {
	return
}

func pages(currentPage string) (int, int) {
	i, err := strconv.Atoi(currentPage)

	if err != nil {
		i = 1
	}

	previous := i - 1
	next := i + 1

	var displayedPrevious int
	if i < 2 {
		displayedPrevious = 1
	} else {
		displayedPrevious = previous
	}

	return displayedPrevious, next
}

func siteRoute(request *http.Request, path string) string {
	tlsStatus := request.TLS
	var urlScheme string

	if tlsStatus == nil {
		urlScheme = "http"
	} else {
		urlScheme = "https"
	}

	var routeURL string

	url, err := url2.Parse(path)
	if err == nil {
		url.Scheme = urlScheme
		url.Host = request.Host

		routeURL = url.String()
	} else {
		routeURL = "#"
	}

	return routeURL
}

func handleFindError(writer http.ResponseWriter, err error) {
	var msg string
	var statusCode int
	if strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
		statusCode = http.StatusNotFound
		msg = "Not Found"
	} else {
		statusCode = http.StatusInternalServerError
		msg = "Server Error"
	}
	log.Println(err)
	writeJSON(writer, ErrorResponse{Message: msg}, statusCode)
}

func handleIndexError(writer http.ResponseWriter, err error) {
	log.Println(err)

	writeJSON(writer, ErrorResponse{
		Message: "Server Error",
	}, 500)
}
