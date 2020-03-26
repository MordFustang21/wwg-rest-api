package main

import (
	"encoding/json"
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
	type Book struct {
		Name   string `json:"name"`
		Author string `json:"author"`
	}

	r.Get("/books/{name}", func(writer http.ResponseWriter, request *http.Request) {
		bookName := chi.URLParam(request, "name")
		b := Book{
			Name:   bookName,
			Author: "Unknown",
		}

		writer.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(b)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})
	// END OMIT

	// SERV_START OMIT
	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}
	// SERV_END OMIT

	return nil
}
