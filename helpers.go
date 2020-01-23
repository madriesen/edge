package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func getJsonFromPostRequest(r *http.Request, v interface{}) error {
	if r.Method != "POST" {
		return errors.New(fmt.Sprintf("expected POST request, received %s", r.Method))
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(bodyBytes, &v)
	if err != nil {
		return err
	}
	return nil
}

func writeSuccessJson(w http.ResponseWriter, v interface{}) {
	var resp struct {
		Value interface{} `json:"value"`
	}
	resp.Value = v
	jsonBytes, _ := json.Marshal(&resp)
	_, _ = io.WriteString(w, string(jsonBytes))
}

func writeErrorJson(w http.ResponseWriter, e error) {
	var resp struct {
		Message string `json:"error"`
	}
	resp.Message = e.Error()
	jsonBytes, _ := json.Marshal(&resp)
	_, _ = io.WriteString(w, string(jsonBytes))
}
