package core

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetcher(t *testing.T) {
	successServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/") {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer successServer.Close()

	failedServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/") {
			w.WriteHeader(http.StatusForbidden)
		}
	}))
	defer failedServer.Close()

	testCases := []struct {
		name   string
		server *httptest.Server
		hasErr bool
	}{
		{"Success", successServer, false},
		{"Failed", failedServer, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := fetcher(tc.server.URL)

			if tc.hasErr && err == nil {
				t.Errorf("Expected an error but got none")
			}

			if !tc.hasErr && err != nil {
				t.Errorf("Did not expect an error but got one: %v", err)
			}
		})
	}
}
