package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"secretservice/keeper"
	"secretservice/keygenerator"
	"strings"
	"testing"
)

func performRequest(method, url string, body io.Reader) *httptest.ResponseRecorder {
	router := getRouter()
	request, _ := http.NewRequest(method, url, body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	return w
}

func TestIndexPageCase(t *testing.T) {
	w := performRequest("GET", "/", nil)
	if w.Code != 200 {
		t.Error("index page is not 200")
	}
}

func TestSaveMessage(t *testing.T) {
	testmessage := "foo"
	postdata := strings.NewReader(fmt.Sprintf("message=%s", testmessage))
	w := performRequest("POST", "/", postdata)
	if w.Code != 200 {
		t.Error("save is not 200")
	}

	key := keygenerator.Key.Create()
	savedmessage, err := keeper.Keep.Get(key)
	if err != nil {
		t.Error("should be nill")
	}
	if savedmessage != testmessage {
		t.Error("not save properly")
	}

	result := w.Result()
	defer result.Body.Close()
	data, _ := ioutil.ReadAll(result.Body)
	if !strings.Contains(string(data), key) {
		t.Error("not right url")
	}

}
