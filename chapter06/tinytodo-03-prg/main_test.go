package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestToDo追加_リダイレクトする(t *testing.T) {
	todoList = []string{}

	form := url.Values{}
	form.Set("todo", "散歩")
	req := httptest.NewRequest("POST", "/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	handleAdd(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Fatalf("expected status %d, got %d", http.StatusSeeOther, rr.Code)
	}
	if location := rr.Header().Get("Location"); location != "/todo" {
		t.Fatalf("expected redirect to /todo, got %q", location)
	}
	if len(todoList) != 1 || todoList[0] != "散歩" {
		t.Fatalf("expected todo list to contain added item")
	}
}
