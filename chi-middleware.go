package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// calling our run function and handling the error
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// run is used so we can just return errors and handle a single exit point in main.
func run() error {
	r := chi.NewRouter()
	// START OMIT
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// pull the authorization header and verify a value has been given
			authValue := r.Header.Get("Authorization")
			if authValue == "" {
				http.Error(w, "authorization required", http.StatusUnauthorized)
				return
			}

			// auth has been provided allow the router to call the route
			next.ServeHTTP(w, r)
		})
	})
	// END OMIT

	// SERV_START OMIT
	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}
	// SERV_END OMIT

	return nil
}
