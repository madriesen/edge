package auth

import (
	"errors"
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

	// check what function is need to be called and run
	success, err = runGrpcHelper(action, success, err, req)

	if helpers.HasError(err, w) {
		return
	}

	helpers.WriteSuccessJson(w, success)
}

func runGrpcHelper(action string, success interface{}, err error, req helpers.AuthRequest) (interface{}, error) {
	switch action {
	case "register":
		return helpers.RunGrpc(service, helpers.GrpcRegister(req))
	case "authenticate":
		return helpers.RunGrpc(service, helpers.GrpcLogin(req))
	case "checkRegister":
		return helpers.RunGrpc(service, helpers.GrpcCheckEmailRegistered(req))
	default:
		return nil, errors.New("no action found")
	}
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
