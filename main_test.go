package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"secretservice/keygenerator"
	"secretservice/storage"
	"strings"
	"testing"
)

const MESSAGEMAXLENGTH = 1024

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
	key, _ := keygenerator.Key.Create()
	_, err := storage.Keep.Get(key)
	if err == nil {
		t.Error("should be nill")
	}
	storage.Keep.Set(key, testmessage, 0)
	savedmessage, _ := storage.Keep.Get(key)
	if savedmessage != testmessage {
		t.Error("not save properly")
	}
}

func TestReadMessage(t *testing.T) {
	testMessage := "hello World!"
	key, _ := keygenerator.Key.Create()
	storage.Keep.Set(key, testMessage, 0)
	router := getRouter()
	request, _ := http.NewRequest("GET", fmt.Sprintf("/%s", key), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	if w.Code != 200 {
		t.Error("Response not 200")
	}
	result := w.Result()
	defer result.Body.Close()
	data, _ := ioutil.ReadAll(result.Body)
	if !strings.Contains(string(data), testMessage) {
		t.Error("result page without key")
	}

	if _, err := storage.Keep.Get(key); err == nil {
		t.Error("Key still in Keeper")
	}
}

func TestReadMessageNotFound(t *testing.T) {
	key, err := keygenerator.Key.Create()
	if err != nil {
		t.Error("must be nill")
	}
	router := getRouter()
	request, _ := http.NewRequest("GET", fmt.Sprintf("/%s", key), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	if w.Code != 404 {
		t.Error("empty message must be 404")
	}
}

func TestOneReaderAtaTime(t *testing.T) {
	testMessage := "hello World!"
	key, _ := keygenerator.Key.Create()
	storage.Keep.Set(key, testMessage, 0)
	router := getRouter()
	resultChannel := make(chan int, 100)

	for i := 0; i < 2; i++ {
		go func(c chan int) {
			request, _ := http.NewRequest("GET", fmt.Sprintf("/%s", key), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, request)
			resultChannel <- w.Code
		}(resultChannel)
	}
	twozz := 0
	for i := 0; i < 2; i++ {
		req := <-resultChannel
		if req == 200 {
			twozz++
		}
	}
	if twozz != 1 {
		t.Error("Should be only 1 ok")
	}
}

func TestBigMessage(t *testing.T) {
	testmessage := strings.Repeat("TEST1234", 1+MESSAGEMAXLENGTH/8)
	postdata := strings.NewReader(fmt.Sprintf("message=%s", testmessage))
	w := performRequest("POST", "/", postdata)
	if w.Code != 400 {
		t.Error("Must be error, because message too long but ", w.Code)
	}
}

func TestMaxTTLVALIDATION(t *testing.T) {
	testmessage := "foo"
	ttl := "999999"
	postdata := strings.NewReader(fmt.Sprintf("message=%s&ttl=%s", testmessage, ttl))
	w := performRequest("POST", "/", postdata)
	if w.Code != 400 {
		t.Error("Must be error, because message too long but ", w.Code)
	}
}
