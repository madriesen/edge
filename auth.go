package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net/http"

	pb "github.com/acubed-tm/edge/protofiles"
)

const authIp = "authenticationms.acubed:50551"

func doRegister(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Token    string `json:"token"`
	}

	err := getJsonFromPostRequest(r, &req)
	if err != nil {
		writeErrorJson(w, err)
		return
	}

	success, err := runGrpc(authIp, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := pb.NewAuthServiceClient(conn)
		resp, err := c.Register(ctx, &pb.RegisterRequest{Email: req.Email, Password: req.Password, VerificationToken: req.Token})
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
		c := pb.NewAuthServiceClient(conn)
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

func doIsEmailRegistered(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	err := getJsonFromPostRequest(r, &req)
	if err != nil {
		writeErrorJson(w, err)
		return
	}

	success, err := runGrpc(authIp, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := pb.NewAuthServiceClient(conn)
		resp, err := c.IsEmailRegistered(ctx, &pb.IsEmailRegisteredRequest{Email: req.Email})
		if err != nil {
			return nil, errors.New(err.Error())
		}
		return resp.IsRegistered, nil
	})

	if err != nil {
		writeErrorJson(w, err)
		return
	}

	writeSuccessJson(w, success)
}
