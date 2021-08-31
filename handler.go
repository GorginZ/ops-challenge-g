package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type handler struct {
	key []byte
	// FIXME not thread safe
	// use not map or protect it?
	stats map[string]uint64
}

type token struct {
	Token []byte `json:"token"`
}
type appMetrics struct {
	Stats map[string]uint64 `json:"stats"`
}

func (h *handler) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (h *handler) token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	var mutex = &sync.Mutex{}

	mutex.Lock()
	h.stats["requests"] += 1
	mutex.Unlock()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		doInternalServerError(w, r, err)
	} else {

		out := createMAC(body, h.key)
		if out == nil {
			doInternalServerError(w, r, err)
			return
		}
		metric := token{Token: out}
		enc.Encode(metric)
		// fmt.Fprintf(w, "%x", out)
		w.WriteHeader(201)
		// enc.Encode(201)
	}
}

func doInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	enc := json.NewEncoder(w)
	fmt.Println("error: ", err)
	w.WriteHeader(500)
	enc.Encode(500)
}

func (h *handler) metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	h.stats["requests"] += 1

	metric := appMetrics{}
	metric.Stats = h.stats

	// FIXME error not checked
	enc.Encode(metric)
	// FIXME error not checked
	w.WriteHeader(200)
}

func createMAC(message, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}
