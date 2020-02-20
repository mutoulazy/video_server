package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/session"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	user := &defs.UserCredential{}

	if err := json.Unmarshal(res, user); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(user.Username, user.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(user.Username)
	signUpResult := &defs.SignedUp{SessionId: id, Success: true}

	if responseJson, err := json.Marshal(signUpResult); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(responseJson), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userName := ps.ByName("user_name")
	io.WriteString(w, "Login User "+userName)
}
