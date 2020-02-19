package main

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpStatusCode)
	resJson, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resJson))
}

func sendNormalResponse(w http.ResponseWriter, resp string, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
	io.WriteString(w, resp)
}
