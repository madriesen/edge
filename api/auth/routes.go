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

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/authenticate", authenticate)
	router.Post("/reauthenticate", reauthenticate)
	router.Post("/register", register)
	router.Post("/check-registration", checkRegistration)
	return router
}
