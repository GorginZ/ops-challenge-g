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

<<<<<<< HEAD
type metrics struct {
	Token []byte `json:"stats"`
=======
type token struct {
	Token []byte `json:"token"`
}
type appMetrics struct {
	Stats map[string]uint64 `json:"stats"`
>>>>>>> 29c20c2e2f0378b74f018c0735a471c48362c3ed
}

func (h *handler) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (h *handler) token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
<<<<<<< HEAD
	h.stats["requests"] += 1
=======
	enc := json.NewEncoder(w)
	var mutex = &sync.Mutex{}

	mutex.Lock()
	h.stats["requests"] += 1
	mutex.Unlock()
>>>>>>> 29c20c2e2f0378b74f018c0735a471c48362c3ed

	body, err := io.ReadAll(r.Body)
	if err != nil {
		doInternalServerError(w, r, err)
	} else {

		out := createMAC(body, h.key)
		if out == nil {
			doInternalServerError(w, r, err)
			return
		}
<<<<<<< HEAD
		metric := metrics{Token: out}
		json.NewEncoder(w).Encode(metric)

		// fmt.Fprintf(w, "%x", out)

		w.WriteHeader(201)

		// enc.Encode(201)
=======
		metric := token{Token: out}
		enc.Encode(metric)
>>>>>>> 29c20c2e2f0378b74f018c0735a471c48362c3ed
	}
}

func doInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	enc := json.NewEncoder(w)
	fmt.Println("error: ", err)
	enc.Encode((err))
	w.WriteHeader(500)
}

func (h *handler) metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
<<<<<<< HEAD
	h.stats["requests"] += 1
=======
>>>>>>> 29c20c2e2f0378b74f018c0735a471c48362c3ed
	enc := json.NewEncoder(w)
	h.stats["requests"] += 1

	metric := appMetrics{}
	metric.Stats = h.stats

	// FIXME error not checked
	enc.Encode(metric)
	// FIXME error not checked
<<<<<<< HEAD
	w.WriteHeader(200)
=======
>>>>>>> 29c20c2e2f0378b74f018c0735a471c48362c3ed
}

func createMAC(message, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}
