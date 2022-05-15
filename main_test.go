package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(method, url string) *httptest.ResponseRecorder {
	router := getRouter()
	request, _ := http.NewRequest(method, url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	return w
}

func TestIndexPageCase(t *testing.T) {
	w := performRequest("GET", "/")
	if w.Code != 200 {
		t.Error("index page is not 200")
	}
}
