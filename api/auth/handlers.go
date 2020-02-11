package auth

import (
	"errors"
	"fmt"
	"github.com/acubed-tm/edge/helpers"
	"github.com/acubed-tm/edge/protofiles"
	"google.golang.org/grpc"
	"net/http"
)

const service = "authenticationms.acubed:50551"

func handler(w http.ResponseWriter, r *http.Request, action string) {
	var req helpers.AuthRequest
	var success interface{}

	err := helpers.GetJsonFromPostRequest(r, &req)
	if err != nil {
		helpers.WriteErrorJson(w, r, err)
		return
	}

	success, err := helpers.RunGrpc(service, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := proto.NewAuthServiceClient(conn)
		resp, err := c.Register(ctx, &proto.RegisterRequest{Email: req.Email, Password: req.Password, VerificationToken: req.Token})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("could not log in: %v", err))
		}
		return resp.Success, nil
	})

	if err != nil {
		helpers.WriteErrorJson(w, r, err)
		return
	}

	helpers.WriteSuccessJson(w, r, success)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := helpers.GetJsonFromPostRequest(r, &req)
	if err != nil {
		helpers.WriteErrorJson(w, r, err)
		return
	}

	type reply struct {
		Token    string `json:"token"`
	}

	resp, err := helpers.RunGrpc(service, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := proto.NewAuthServiceClient(conn)
		resp, err := c.Login(ctx, &proto.LoginRequest{Email: req.Email, Password: req.Password})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("could not log in: %v", err))
		}
		return reply{Token: resp.Token}, nil
	})

	if err != nil {
		helpers.WriteErrorJson(w, r, err)
		return
	}

	helpers.WriteSuccessJson(w, r, resp)
}

func checkRegistration(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	err := helpers.GetJsonFromPostRequest(r, &req)
	if err != nil {
		helpers.WriteErrorJson(w, r, err)
		return
	}

	success, err := helpers.RunGrpc(service, func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := proto.NewAuthServiceClient(conn)
		resp, err := c.IsEmailRegistered(ctx, &proto.IsEmailRegisteredRequest{Email: req.Email})
		if err != nil {
			return nil, errors.New(err.Error())
		}
		return resp.IsRegistered, nil
	})

	if err != nil {
		helpers.WriteErrorJson(w, r, err)
		return
	}

	helpers.WriteSuccessJson(w, r, success)
}
