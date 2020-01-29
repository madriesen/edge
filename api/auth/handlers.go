package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/acubed-tm/edge/helpers"

	"google.golang.org/grpc"

	pb "github.com/acubed-tm/edge/protofiles"
)

const service = "authenticationms.acubed:50551"

func register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Token    string `json:"token"`
	}

	err := helpers.GetJsonFromPostRequest(r, &req)
	if err != nil {
		helpers.WriteErrorJson(w, err)
		return
	}

	success, err := helpers.RunGrpc(service, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := pb.NewAuthServiceClient(conn)
		resp, err := c.Register(ctx, &pb.RegisterRequest{Email: req.Email, Password: req.Password, VerificationToken: req.Token})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("could not log in: %v", err))
		}
		return resp.Success, nil
	})

	if err != nil {
		helpers.WriteErrorJson(w, err)
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
	if err != nil {
		helpers.WriteErrorJson(w, err)
		return
	}

	success, err := helpers.RunGrpc(service, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := pb.NewAuthServiceClient(conn)
		resp, err := c.Login(ctx, &pb.LoginRequest{Email: req.Email, Password: req.Password})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("could not log in: %v", err))
		}
		return resp.Success, nil
	})

	if err != nil {
		helpers.WriteErrorJson(w, err)
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
	if err != nil {
		helpers.WriteErrorJson(w, err)
		return
	}

	success, err := helpers.RunGrpc(service, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := pb.NewAuthServiceClient(conn)
		resp, err := c.Login(ctx, &pb.LoginRequest{Email: req.Email, Password: req.Password})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("could not log in: %v", err))
		}
		return resp.Success, nil
	})

	if err != nil {
		helpers.WriteErrorJson(w, err)
		return
	}

	helpers.WriteSuccessJson(w, success)
}

func checkRegistration(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	err := helpers.GetJsonFromPostRequest(r, &req)
	if err != nil {
		helpers.WriteErrorJson(w, err)
		return
	}

	success, err := helpers.RunGrpc(service, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := pb.NewAuthServiceClient(conn)
		resp, err := c.IsEmailRegistered(ctx, &pb.IsEmailRegisteredRequest{Email: req.Email})
		if err != nil {
			return nil, errors.New(err.Error())
		}
		return resp.IsRegistered, nil
	})

	if err != nil {
		helpers.WriteErrorJson(w, err)
		return
	}

	helpers.WriteSuccessJson(w, success)
}
