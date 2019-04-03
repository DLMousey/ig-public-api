package main

import (
	"fmt"
	"net/http"
)

/* contains all the routes */

func (a *app) setRoutes() {
	/*	sample code
		a.router.HandleFunc("/api/", s.handleAPI())
		a.router.HandleFunc("/api/:users", s.handleUsers())
		a.router.HandleFunc("/api/:images", s.handleImages())
	*/
	a.router.HandleFunc("/test", a.handleTest).Methods("GET")
}

func (a *app) handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route is working")
}

/* Stole this code somewhere and it looks clean asf
// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
*/
