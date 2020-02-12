package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	io.WriteString(w, "Create User Handler!!!")
}