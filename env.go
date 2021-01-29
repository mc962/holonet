package main

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"holonet/data"
	"holonet/web"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const DefaultPort = 8080

type Env struct {
	Router *mux.Router
	Wait   time.Duration
	Conn   *pgx.Conn
}

func (env *Env) Initialize() {
	env.Conn = data.Initialize()
	env.Router = web.Initialize()
}

func (env *Env) Run(port string) {
	var appPort int
	if len(port) > 0 {
		var err error
		appPort, err = strconv.Atoi(port)
		if err != nil {
			log.Printf("Invalid port %s\n", port)
			os.Exit(1)
		}
	} else {
		appPort = DefaultPort
	}

	loggedRouter := handlers.LoggingHandler(os.Stdout, env.Router)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", appPort),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      loggedRouter,
	}

	log.Printf("Starting at port %d\n", appPort)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	channel := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(channel, os.Interrupt)

	// Block until we receive our signal.
	<-channel

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), env.Wait)

	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	// Error may be ignored, as we will be existing anyway
	_ = srv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("\n\nShutting down...")
	os.Exit(0)
}
