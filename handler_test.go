package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
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
		req, _ := http.NewRequest("POST", "/", strings.NewReader(tt.body))

		h.token(rec, req)

		mac := createMAC([]byte(tt.body), h.key)

		var tok []byte
		err := json.Unmarshal([]byte(rec.Body.Bytes()), &tok)
		if err != nil {
			panic(err)
		}
		actual := tok

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
