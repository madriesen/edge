package main

import (
	"io"
	"net/http"
)

func DoLogin(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "this is a login endpoint")
}
