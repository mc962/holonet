package web

import (
	"net/http"
)

func RootHandler(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, map[string]string{
		"version": "0.1.0",
	}, 200)
}

func APIRootHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Accept", "application/json")

	siteRoutes := map[string]string{
		"people":    siteRoute(request, "/api/people/"),
		"planets":   siteRoute(request, "/api/planets/"),
		"films":     siteRoute(request, "/api/films/"),
		"species":   siteRoute(request, "/api/species/"),
		"vehicles":  siteRoute(request, "/api/vehicles/"),
		"starships": siteRoute(request, "/api/starships/"),
	}

	writeJSON(writer, siteRoutes, 200)
}
