package helpers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

func GetJsonFromPostRequest(r *http.Request, v interface{}) error {
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

func WriteSuccessJson(w http.ResponseWriter, v interface{}) {
	log.Printf("Returning success: %v", v)
	var resp struct {
		Value interface{} `json:"value"`
	}
	resp.Value = v
	jsonBytes, _ := json.Marshal(&resp)
	_, _ = io.WriteString(w, string(jsonBytes))
}

func writeErrorJson(w http.ResponseWriter, e error) {
	log.Printf("Returning error: %v", e.Error())
	var resp struct {
		Message string `json:"error"`
	}
	resp.Message = e.Error()
	jsonBytes, _ := json.Marshal(&resp)
	_, _ = io.WriteString(w, string(jsonBytes))
}

func RunGrpc(ip string, f func(context.Context, *grpc.ClientConn) (interface{}, error)) (interface{}, error) {
	log.Printf("Starting gRPC connection to %s", ip)
	conn, err := grpc.Dial(ip, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("did not connect: %v", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	ret, err := f(ctx, conn)

	cancel()
	_ = conn.Close()

	return ret, err
}

func HasError(err error, w http.ResponseWriter) bool {
	if err != nil {
		writeErrorJson(w, err)
		return true
	}
	return false
}
