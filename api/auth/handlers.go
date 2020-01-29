package auth

import (
	"net/http"

	"github.com/acubed-tm/edge/helpers"
)

const service = "authenticationms.acubed:50551"

func handler(w http.ResponseWriter, r *http.Request, action string) {
	var req helpers.AuthRequest
	var success interface{}

	err := helpers.GetJsonFromPostRequest(r, &req)
	if helpers.HasError(err, w) {
		return
	}

	// check what function is need to be called
	success, err = runGrpcHelper(action, success, err, req)

	if helpers.HasError(err, w) {
		return
	}

	// create token and return

	helpers.WriteSuccessJson(w, success)
}

func runGrpcHelper(action string, success interface{}, err error, req helpers.AuthRequest) (interface{}, error) {
	switch action {
	case "register":
		success, err = helpers.RunGrpc(service, helpers.GrpcRegister(req))
	case "authenticate":
		success, err = helpers.RunGrpc(service, helpers.GrpcLogin(req))
	case "checkRegister":
		success, err = helpers.RunGrpc(service, helpers.GrpcCheckEmailRegistered(req))
	}
	return success, err
}

func register(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "register")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "authenticate")
}

func reauthenticate(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "authenticate")
}

func checkRegistration(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "checkRegister")
}

