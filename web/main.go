package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", HomeHandler)

	router.POST("/", HomeHandler)

	// router.GET("/userhome", userHomeHandler)

	// router.POST("/userhome", userHomeHandler)

	// router.POST("/api", apiHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./template"))
	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe("localhost:8080", r)
}
