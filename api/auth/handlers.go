package auth

import (
	"net/http"

	"github.com/acubed-tm/edge/helpers"
)

const service = "authenticationms.acubed:50551"

func register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Token    string `json:"token"`
	}

	err := helpers.GetJsonFromPostRequest(r, &req)
	if helpers.HasError(err, w) {
		return
	}

	success, err := helpers.RunGrpc(service, helpers.GrpcRegister(req))

	if helpers.HasError(err, w) {
		return
	}

	helpers.WriteSuccessJson(w, success)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := helpers.GetJsonFromPostRequest(r, &req)
	if helpers.HasError(err, w) {
		return
	}

	success, err := helpers.RunGrpc(service, helpers.GrpcLogin(req))

	if helpers.HasError(err, w) {
		return
	}

	helpers.WriteSuccessJson(w, success)
}

func reauthenticate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := helpers.GetJsonFromPostRequest(r, &req)
	if helpers.HasError(err, w) {
		return
	}

	success, err := helpers.RunGrpc(service, helpers.GrpcLogin(req))

	if helpers.HasError(err, w) {
		return
	}

	helpers.WriteSuccessJson(w, success)
}

func checkRegistration(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	err := helpers.GetJsonFromPostRequest(r, &req)
	if helpers.HasError(err, w) {
		return
	}

	success, err := helpers.RunGrpc(service, helpers.GrpcCheckEmailRegistered(req))

	if helpers.HasError(err, w) {
		return
	}

	helpers.WriteSuccessJson(w, success)
}

