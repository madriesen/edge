package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/acubed-tm/edge/helpers"
	proto "github.com/acubed-tm/edge/protofiles"
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"google.golang.org/grpc"
	"net/http"
	"reflect"
	"time"
)

func clientInvoker() (interface {}, context.Context, context.CancelFunc, error) {
	const address = "authenticationms.acubed:50551"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, nil, errors.New(fmt.Sprintf("did not connect: %v", err))
	}

	client := proto.NewAuthServiceClient(*conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	return client, ctx, cancel, nil
}

func handle(
	writer http.ResponseWriter,
	request *http.Request,
	clientInvoker func() (interface {}, context.Context, context.CancelFunc, error),
	action string) (interface{}, error) {

	client, ctx, cancel, err := clientInvoker()

	method := reflect.ValueOf(client).MethodByName(action)

	r := method.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.New(method.Type().In(1)),
	})

	if !r[1].IsNil() {
		return nil, errors.New(fmt.Sprintf("could not log in: %v", r[1].Interface()))
	}

	panic("todo")
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/authenticate", authenticate)
	router.Post("/reauthenticate", reauthenticate)
	router.Post("/register", register)
	router.Post("/check-registration", checkRegistration)
	return router
}
