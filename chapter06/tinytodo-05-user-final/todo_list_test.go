package main

import "testing"

func TestToDoリスト追加_要素が増える(t *testing.T) {
	list := NewToDoList()
	list.Append("買い物")
	list.Append("掃除")

	if len(list.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(list.Items))
	}
	if list.Items[0] != "買い物" || list.Items[1] != "掃除" {
		t.Fatalf("unexpected items: %v", list.Items)
	}
}
