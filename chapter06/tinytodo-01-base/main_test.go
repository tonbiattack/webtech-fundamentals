package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestToDo表示_一覧が含まれる(t *testing.T) {
	todoList = []string{"顔を洗う", "朝食を食べる"}

	req := httptest.NewRequest("GET", "/todo", nil)
	rr := httptest.NewRecorder()

	handleTodo(rr, req)

	body := rr.Body.String()
	for _, item := range todoList {
		if !strings.Contains(body, item) {
			t.Fatalf("expected response to contain %q", item)
		}
	}
}
