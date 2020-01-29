package auth

import (
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/authenticate", authenticate)
	router.Post("/reauthenticate", reauthenticate)
	router.Post("/register", register)
	router.Post("/check-registration", checkRegistration)
	return router
}
