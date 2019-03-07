package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", strings.NewReader(""))
	rr := httptest.NewRecorder()
	handler(rr, req)

	out, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}

	want := "Hello World!"
	if got := string(out); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
