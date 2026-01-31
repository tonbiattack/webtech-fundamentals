package main

import (
	"flag"
	"os"
	"testing"
)

func resetFlags(t *testing.T) {
	t.Helper()
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}

func Testポート番号取得_環境変数優先(t *testing.T) {
	resetFlags(t)
	t.Setenv("PORT", "9090")

	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs })
	os.Args = []string{"cmd", "-p", "7070"}

	if got := getPortNumber(); got != 9090 {
		t.Fatalf("expected 9090, got %d", got)
	}
}

func Testポート番号取得_フラグ指定(t *testing.T) {
	resetFlags(t)
	t.Setenv("PORT", "")

	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs })
	os.Args = []string{"cmd", "-p", "7070"}

	if got := getPortNumber(); got != 7070 {
		t.Fatalf("expected 7070, got %d", got)
	}
}
