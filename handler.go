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
	key   []byte
	stats map[string]uint64
	lock  sync.Mutex
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
	h.lock.Lock()
	defer h.lock.Unlock()
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
		metric := token{Token: out}
		enc.Encode(metric)
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
	enc := json.NewEncoder(w)

	metric := appMetrics{}
	metric.Stats = h.stats

	// FIXME error not checked
	enc.Encode(metric)
	// FIXME error not checked
}

func createMAC(message, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}
