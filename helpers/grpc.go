package helpers

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/acubed-tm/edge/protofiles"
	"google.golang.org/grpc"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func GrpcRegister(req AuthRequest) func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
	return func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := pb.NewAuthServiceClient(conn)
		resp, err := c.Register(ctx, &pb.RegisterRequest{Email: req.Email, Password: req.Password, VerificationToken: req.Token})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("could not log in: %v", err))
		}
		return resp.Success, nil
	}
}


func GrpcLogin(req AuthRequest) func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
	return func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := pb.NewAuthServiceClient(conn)
		resp, err := c.Login(ctx, &pb.LoginRequest{Email: req.Email, Password: req.Password})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("could not log in: %v", err))
		}
		return resp.Success, nil
	}
}


func GrpcCheckEmailRegistered(req AuthRequest) func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
	return func(ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
		// Contact the server and print out its response.
		c := pb.NewAuthServiceClient(conn)
		resp, err := c.IsEmailRegistered(ctx, &pb.IsEmailRegisteredRequest{Email: req.Email})
		if err != nil {
			return nil, errors.New(err.Error())
		}
		return resp.IsRegistered, nil
	}
}