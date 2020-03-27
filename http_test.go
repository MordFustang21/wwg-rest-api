package main_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("world"))
	})
}

// SERVER OMIT
func Test_Server(t *testing.T) {
	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}
	// this can leak so it needs to be closed
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "world" {
		t.Fatalf("expected %s got %s", "world", body)
	}
}
// SERVER OMIT

// RECORDER OMIT
func Test_ResponseRecorder(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("world"))
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Fatal("got unexpected status code", resp.StatusCode)
	}

	if string(body) != "world" {
		t.Fatalf("response body expected %s got %s", "world", body)
	}
}
// RECORDER OMIT