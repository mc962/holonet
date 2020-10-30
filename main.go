package main

import (
	"log"
	"os"
)

func main() {
	// create a new application instance, and run through application initialization steps
	app := Env{}
	app.Initialize()

	rows, err := app.Conn.Query("SELECT NOW();")

	if err != nil {
		log.Fatalf("Fatal err %s", err)
	}

	defer rows.Close()

	// run application (server)
	app.Run(os.Getenv("PORT"))
}
