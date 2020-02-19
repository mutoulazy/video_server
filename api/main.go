package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	return router
}

func main() {
	router := RegisterHandlers()
	middleWareHandler := NewMiddleWareHandler(router)
	http.ListenAndServe(":8080", middleWareHandler)
}
