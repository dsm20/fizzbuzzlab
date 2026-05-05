package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		query      string
		wantCode   int
		wantBody   string
	}{
		{"?n=1", http.StatusOK, "1"},
		{"?n=3", http.StatusOK, "Fizz"},
		{"?n=5", http.StatusOK, "Buzz"},
		{"?n=15", http.StatusOK, "FizzBuzz"},
		{"?n=abc", http.StatusBadRequest, ""},
		{"", http.StatusBadRequest, ""},
	}

	for _, tc := range tests {
		req := httptest.NewRequest("GET", "/"+tc.query, nil)
		w := httptest.NewRecorder()
		handler(w, req)

		if w.Code != tc.wantCode {
			t.Errorf("GET /%s: status = %d, want %d", tc.query, w.Code, tc.wantCode)
		}
		if tc.wantBody != "" && !strings.Contains(w.Body.String(), tc.wantBody) {
			t.Errorf("GET /%s: body = %q, want %q", tc.query, w.Body.String(), tc.wantBody)
		}
	}
}
