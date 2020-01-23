package main

import (
	"io"
	"net/http"
)

func doLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := getJsonFromPostRequest(r, &req)
	if err != nil {
		writeErrorJson(w, err)
		return
	}

	// TODO
	// writeSuccessJson()

	_, _ = io.WriteString(w, "nice")
}
