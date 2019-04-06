/*
Contains all the routes
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/nacleric/ig-public-api/common"
)

//mock Data
type mockUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func (a *app) setRoutes() {
	/*	sample routes
		a.router.HandleFunc("/api/", s.handleAPI()).Methods("GET")
		a.router.HandleFunc("/api/:users", s.handleUsers()).Methods("GET")
		a.router.HandleFunc("/api/:images", s.handleImages()).Methods("GET")
	*/
	a.Get("/test", a.handleTest)
}

/* HANDLERS */

// test handler
func (a *app) handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route is working")

	foo := &mockUser{Firstname: "Bob", Lastname: "Ross", Age: 53}
	common.RespondJSON(w, http.StatusCreated, foo)
}

// Get wraps the router for GET method
func (a *app) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("GET")
}
