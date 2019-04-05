/*contains all the routes*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//mock Data
type mockUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func (a *app) setRoutes() {
	/*	sample routes
		a.router.HandleFunc("/api/", s.handleAPI())
		a.router.HandleFunc("/api/:users", s.handleUsers())
		a.router.HandleFunc("/api/:images", s.handleImages())
	*/
	a.Get("/test", a.handleTest)
}

/* HANDLERS */

func (a *app) handleTest(w http.ResponseWriter, r *http.Request) {
	foo := &mockUser{Firstname: "Bob", Lastname: "Ross", Age: 53}
	fmt.Println("route is working")

	// maps datatype to json
	response, err := json.Marshal(foo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// Get wraps the router for GET method
func (a *app) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("GET")
}
