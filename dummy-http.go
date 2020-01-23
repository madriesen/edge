package main

import (
	"io"
	"log"
	"net/http"
)

func startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":8080"}

	// routes
	http.HandleFunc("/hello", func (w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hellom world!")
	})

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	// returning reference so caller can call Shutdown()
	return srv
}

func main() {
	log.Printf("starting HTTP server")
	_ = startHttpServer()
	// sleep forever
	select {}
}