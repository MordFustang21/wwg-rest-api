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
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			writer.Write([]byte("world"))
		default:
			http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
	// END OMIT

	// SERV_START OMIT
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}
	// SERV_END OMIT

	return nil
}
