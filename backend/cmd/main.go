package main

import (
	"log"
	"net/http"

	"github.com/chtozamm/new-project/backend/internal/monitoring"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", monitoring.PrometheusHandler())
	mux.HandleFunc("/", rootHandler)

	muxWithMiddleware := monitoring.PrometheusMiddleware(mux)

	if err := http.ListenAndServe("0.0.0.0:4000", muxWithMiddleware); err != nil {
		log.Fatalln(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Hello there!"}`))
}
