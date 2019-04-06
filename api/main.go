/*
Main File
*/
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type app struct {
	// db *db
	router *mux.Router
}

// initializes the routes
func (a *app) Initialize() {
	a.router = mux.NewRouter()
	a.setRoutes()
	fmt.Println("api is running")
	log.Fatal(http.ListenAndServe(":8000", a.router))
}

func main() {
	var app app
	app.Initialize()
}
