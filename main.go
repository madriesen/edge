package main

import (
	"log"
	"net/http"

	"github.com/acubed-tm/edge/api/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func ShowAPIInfo(w http.ResponseWriter, r *http.Request) {
	type ServerInfo struct {
		Message       string `json:"message"`
		LatestVersion string `json:"version"`
	}
	info := ServerInfo{
		Message:       "Service available.",
		LatestVersion: "v1",
	}
	render.JSON(w, r, info) // A chi router helper for serializing and returning json
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Get("/", ShowAPIInfo)
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/auth", auth.Routes())
	})

	return router
}

func main() {
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":8081", router))
}
