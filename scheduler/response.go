package main

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, statusCode int, resp string) {
	w.WriteHeader(statusCode)
	io.WriteString(w, resp)
}
