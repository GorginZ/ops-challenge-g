package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type handler struct {
	key []byte
	// FIXME not thread safe
	stats map[string]uint64
}

type metrics struct {
	Token []byte `json:"token"`
}

func (h *handler) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (h *handler) token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	h.stats["requests"] += 1

	body, err := io.ReadAll(r.Body)
	if err != nil {
		doInternalServerError(w, r, err)
	} else {

		out := createMAC(body, h.key)
		if out == nil {
			doInternalServerError(w, r, err)
			return
		}
		metric := metrics{Token: out}
		json.NewEncoder(w).Encode(metric)

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
	h.stats["requests"] += 1
	enc := json.NewEncoder(w)
	// FIXME error not checked
	enc.Encode(h.stats)
	// FIXME error not checked
	w.WriteHeader(200)
}

func createMAC(message, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}
