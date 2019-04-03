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
	a.router.HandleFunc("/test", handleTest).Methods("GET")
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route is working")
}
