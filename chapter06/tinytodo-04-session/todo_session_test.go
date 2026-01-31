package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestToDoリスト取得_初期化される(t *testing.T) {
	todoLists = make(map[string][]string)

	list := getTodoList("session-1")
	if len(list) != 0 {
		t.Fatalf("expected empty todo list")
	}
	if _, ok := todoLists["session-1"]; !ok {
		t.Fatalf("expected todo list to be created")
	}
}

func Testセッション開始_クッキー設定(t *testing.T) {
	req := httptest.NewRequest("GET", "/todo", nil)
	rr := httptest.NewRecorder()

	sessionId, err := ensureSession(rr, req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sessionId == "" {
		t.Fatalf("expected session id to be set")
	}

	cookies := rr.Result().Cookies()
	found := false
	for _, c := range cookies {
		if c.Name == cookieSessionId {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected session cookie to be set")
	}
}

func TestToDo追加_トリムとエスケープ(t *testing.T) {
	todoLists = make(map[string][]string)
	todoLists["sid"] = []string{}

	form := url.Values{}
	form.Set("todo", "  <b>abc</b>  ")
	req := httptest.NewRequest("POST", "/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(&http.Cookie{Name: cookieSessionId, Value: "sid"})
	rr := httptest.NewRecorder()

	handleAdd(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Fatalf("expected status %d, got %d", http.StatusSeeOther, rr.Code)
	}
	got := todoLists["sid"]
	if len(got) != 1 || got[0] != "&lt;b&gt;abc&lt;/b&gt;" {
		t.Fatalf("unexpected todo list: %v", got)
	}
}
