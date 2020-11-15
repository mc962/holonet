package main

import (
	"os"
)

func main() {
	// create a new application instance, and run through application initialization steps
	app := Env{}
	app.Initialize()

	// run application (server)
	app.Run(os.Getenv("PORT"))
}
