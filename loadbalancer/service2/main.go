package main

import (
	"fmt"
	"net/http"
)

func main() {
	if err := mainErr(); err != nil {
		fmt.Println("error ", err)
	}
}

type server struct{}

func mainErr() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("healthcheck for server 2")

		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is server 2"))
	})

	if err := http.ListenAndServe(":8081", mux); err != nil {
		return err
	}

	return nil
}
