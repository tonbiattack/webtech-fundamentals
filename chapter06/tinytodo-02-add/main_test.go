package main

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestToDo追加_一覧に反映(t *testing.T) {
	todoList = []string{}

	form := url.Values{}
	form.Set("todo", "買い物")
	req := httptest.NewRequest("POST", "/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	handleAdd(rr, req)

	if len(todoList) != 1 || todoList[0] != "買い物" {
		t.Fatalf("expected todo list to contain added item")
	}
}
