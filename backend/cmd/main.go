package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello there!"}`))
	})

	if err := http.ListenAndServe("0.0.0.0:4000", mux); err != nil {
		log.Fatalln(err)
	}
}
