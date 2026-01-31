package main

import (
	"net/http/httptest"
	"testing"
)

func TestHello_レスポンス本文(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	hello(rr, req)

	if got := rr.Body.String(); got != "Hello, Web application!" {
		t.Fatalf("expected %q, got %q", "Hello, Web application!", got)
	}
}
