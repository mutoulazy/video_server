package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	return router
}

func main() {
	router := RegisterHandlers()
	http.ListenAndServe(":8080", router)
}
