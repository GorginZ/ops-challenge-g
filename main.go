package main

import (
	"log"
	"net/http"
	"os"
)

var requestIPs []string

func main() {
	router()
}

func router() {
	h := &handler{
		key:   []byte(os.Getenv("SECRET")),
		stats: map[string]uint64{"requests": 0},
	}
	http.HandleFunc("/token", h.token)
	http.HandleFunc("/metrics", h.metrics)
	http.HandleFunc("/health", h.health)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
