package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestToken(t *testing.T) {
	tests := []struct {
		body string
	}{
		{"123"},
	}

	h := handler{
		stats: make(map[string]uint64),
		key:   []byte("some-baked-in-secret"),
	}

	for _, tt := range tests {
		rec := httptest.NewRecorder()
		// FIXME testing wrong method
		req, _ := http.NewRequest("POST", "/", strings.NewReader(tt.body))

		h.token(rec, req)

		mac := createMAC([]byte(tt.body), h.key)
		actual, _ := hex.DecodeString(rec.Body.String())

		if !hmac.Equal(actual, mac) {
			t.Errorf("failed to validate hmac")
		}

	}
}

func validMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
