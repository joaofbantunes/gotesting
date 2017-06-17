package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App : represents the running application
type App struct {
	Router *mux.Router
}

// Initialize : prepares the application for execution
func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

// Run : starts the application
func (app *App) Run() {
	log.Print("Listening...")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/", Logger(index)).Methods("GET")
	app.Router.HandleFunc("/bye", Logger(bye)).Methods("GET")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func bye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bye!")
}
