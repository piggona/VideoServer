package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", HomeHandler)

	router.POST("/", HomeHandler)

	router.GET("/userhome", userHomeHandler)

	router.POST("/userhome", userHomeHandler)

	router.POST("/api", apiHandler)
	
	router.POST("/upload/:vid-id", proxyHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))
	// 访问127.0.0.1:8080/statics/xxx，就会读取./templates/xxx文件
	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe("localhost:8080", r)
}
