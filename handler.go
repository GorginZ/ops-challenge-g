package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

type handler struct {
	key   []byte
	stats map[string]uint64
	lock  sync.Mutex
}

func (h *handler) health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(200)
}

func (h *handler) token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	h.lock.Lock()
	defer h.lock.Unlock()
	h.stats["requests"] += 1
	if r.Method != "POST" {
		w.WriteHeader(400)
		return
	}
	h.giveToken(w, r)
}

func (h *handler) giveToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	// close body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		enc.Encode(err)
		return
	}
	out := createMAC(body, h.key)
	if out != nil {
		w.WriteHeader(200)
		enc.Encode(out)
		return
	}
	w.WriteHeader(500)
}

func createMAC(message, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}

func (h *handler) metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(400)
		return
	}
	h.lock.Lock()
	defer h.lock.Unlock()
	//look for the key!!!
	stats := h.stats
	if stats != nil {
		enc := json.NewEncoder(w)
		enc.Encode(stats)
	} else {
		w.WriteHeader(500)
	}
}
