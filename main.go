package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func logInstance(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Fetching metrics for instance: %s\n", os.Getenv("CF_INSTANCE_INDEX"))
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/metrics", logInstance(promhttp.Handler()))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
