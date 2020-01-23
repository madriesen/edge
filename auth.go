package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net/http"

	pb "github.com/acubed-tm/edge/protofiles"
)

// TODO: don't hardcode, use service discovery
const authIp = "51.136.77.127:50551"

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

	success, err := runGrpc(authIp, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := pb.NewLoginServiceClient(conn)
		resp, err := c.Login(ctx, &pb.LoginRequest{Email: req.Email, Password: req.Password})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("could not log in: %v", err))
		}
		return resp.Success, nil
	})

	if err != nil {
		writeErrorJson(w, err)
		return
	}

	writeSuccessJson(w, success)
}
