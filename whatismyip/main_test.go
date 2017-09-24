package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var payload = `{"ip":"10.0.0.1","ip_decimal":342342342,"country":"Czechia","city":"Prague","hostname":"internet-1231.com"}`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, payload)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestIPinfo(t *testing.T) {
	t.Parallel()
	server := mockServer()
	defer server.Close()

	info := getIPinfo(server.URL)

	expectedInfo := ipInfo{
		IP:       "10.0.0.1",
		Country:  "Czechia",
		City:     "Prague",
		Hostname: "internet-1231.com",
	}

	if info != expectedInfo {
		t.Errorf("Expected %s but got %s ", expectedInfo, info)
	}
}
