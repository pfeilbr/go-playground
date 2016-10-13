package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloReponse(t *testing.T) {
	expect := "hello"

	var helloServer hello
	server := httptest.NewServer(helloServer)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	defer resp.Body.Close()

	var b []byte
	if b, err = ioutil.ReadAll(resp.Body); err != nil {
		t.Errorf("reading reponse body: %v, want %v", err, expect)
	}
	if string(b) != expect {
		t.Errorf("request body mismatch: got %q, want %q", string(b), expect)
	}

}
