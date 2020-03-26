package main

import (
	"log"
	"net/http"
)

func main() {
	// calling our run function and handling the error
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// run is used so we can just return errors and handle a single exit point in main.
func run() error {
	// START OMIT
	customMux := http.NewServeMux()
	customMux.HandleFunc("/hello/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("world"))
	})

	if err := http.ListenAndServe(":8080", customMux); err != nil {
		return err
	}
	// END OMIT

	return nil
}
