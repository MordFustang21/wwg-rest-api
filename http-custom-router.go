package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// START OMIT
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return http.ListenAndServe(":3000", r)
	// END OMIT
}
