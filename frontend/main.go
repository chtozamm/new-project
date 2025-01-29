package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.Handle("/backend", backendHandler())

	if err := http.ListenAndServe("0.0.0.0:3000", mux); err != nil {
		log.Fatalln(err)
	}
}

func backendHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := os.Getenv("BACKEND_URL")
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Failed to make GET request to the backend:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		msg := struct {
			Message string `json:"message"`
		}{}
		json.NewDecoder(resp.Body).Decode(&msg)

		w.Write([]byte(msg.Message))
	})
}
