package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestFileServer_HelloHtml(t *testing.T) {
	expected, err := os.ReadFile("static/hello.html")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	handler := http.FileServer(http.Dir("static"))
	req := httptest.NewRequest("GET", "/hello.html", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
	got := strings.TrimSpace(rr.Body.String())
	want := strings.TrimSpace(string(expected))
	if got != want {
		t.Fatalf("response body mismatch")
	}
}
