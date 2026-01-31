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
	if len(account.ToDoList) != 0 {
		t.Fatalf("expected empty todo list")
	}
}

func Testユーザアカウント作成_不正ID(t *testing.T) {
	m := NewUserAccountManager()
	_, err := m.NewUserAccount("bad id", "secret")
	if err != ErrInvalidUserIdFormat {
		t.Fatalf("expected ErrInvalidUserIdFormat, got %v", err)
	}
}

func Testユーザアカウント作成_重複(t *testing.T) {
	m := NewUserAccountManager()
	_, err := m.NewUserAccount("user1", "secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = m.NewUserAccount("user1", "secret")
	if err != ErrUserAlreadyExists {
		t.Fatalf("expected ErrUserAlreadyExists, got %v", err)
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

func Test有効期限文字列_フォーマット(t *testing.T) {
	expires := time.Date(2025, 1, 2, 3, 4, 5, 0, time.UTC)
	account := NewUserAccount("user1", "secret", expires)
	if got := account.ExpiresText(); got != "2025/01/02 03:04:05" {
		t.Fatalf("unexpected expires text: %s", got)
	}
}
