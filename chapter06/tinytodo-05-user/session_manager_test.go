package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Testセッション開始_クッキー設定(t *testing.T) {
	mgr := NewHttpSessionManager()
	rr := httptest.NewRecorder()

	session, err := mgr.StartSession(rr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if session == nil || session.SessionId == "" {
		t.Fatalf("expected session to be created")
	}

	cookies := rr.Result().Cookies()
	found := false
	for _, c := range cookies {
		if c.Name == CookieNameSessionId {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected session cookie to be set")
	}
}

func Testセッション取得_有効(t *testing.T) {
	mgr := NewHttpSessionManager()
	rr := httptest.NewRecorder()

	session, err := mgr.StartSession(rr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: CookieNameSessionId, Value: session.SessionId})
	got, err := mgr.GetValidSession(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.SessionId != session.SessionId {
		t.Fatalf("expected same session id")
	}
}

func Testセッション取得_期限切れ(t *testing.T) {
	mgr := NewHttpSessionManager()
	expired := NewHttpSession("expired", -time.Minute)
	mgr.sessions["expired"] = expired

	_, err := mgr.getSession("expired")
	if err != ErrSessionExpired {
		t.Fatalf("expected ErrSessionExpired, got %v", err)
	}
	if _, exists := mgr.sessions["expired"]; exists {
		t.Fatalf("expected expired session to be removed")
	}
}
