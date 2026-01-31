package main

import (
	"strings"
	"testing"
	"time"
)

func TestユーザID検証_許可と拒否(t *testing.T) {
	m := NewUserAccountManager()
	if !m.ValidateUserId("user@example.com") {
		t.Fatalf("expected valid user id")
	}
	if m.ValidateUserId("bad id") {
		t.Fatalf("expected invalid user id")
	}
}

func Testユーザアカウント作成_正常系(t *testing.T) {
	m := NewUserAccountManager()
	account, err := m.NewUserAccount("user1", "secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if account.Id != "user1" {
		t.Fatalf("expected user id to be set")
	}
	if account.HashedPassword == "" {
		t.Fatalf("expected hashed password to be set")
	}
	if time.Now().After(account.Expires) {
		t.Fatalf("expected account to be valid in the future")
	}
	if account.ToDoList == nil {
		t.Fatalf("expected todo list to be set")
	}
}

func Test認証_成功と失敗(t *testing.T) {
	m := NewUserAccountManager()
	_, err := m.NewUserAccount("user1", "secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := m.Authenticate("user1", "secret"); err != nil {
		t.Fatalf("expected authentication success, got %v", err)
	}
	if _, err := m.Authenticate("user1", "wrong"); err != ErrLoginFailed {
		t.Fatalf("expected ErrLoginFailed, got %v", err)
	}
}

func Testパスワード生成_長さと文字種(t *testing.T) {
	pwd := MakePassword()
	if len(pwd) != PasswordLength {
		t.Fatalf("expected length %d, got %d", PasswordLength, len(pwd))
	}
	for _, r := range pwd {
		if !strings.ContainsRune(PasswordChars, r) {
			t.Fatalf("unexpected character: %q", r)
		}
	}
}
