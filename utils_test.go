package main

import (
	"testing"
)

func TestHeap(t *testing.T) {
	h := InitMinHeap()

	h.HeapPush(0.1, "Hey1")
	h.HeapPush(3.5, "Hey2")
	h.HeapPush(1.3, "Hey3")
	h.HeapPush(1.0, "Hey4")

	_ = h.SmallestN(2)

	for h.Len() > 0 {
		_, _ = h.HeapPop()
	}
}

func TestMap(t *testing.T) {
	l := []string{"a", "c", "c", "c", "a", "b"}
	m := Occurences(l)
	k := FindMaxInMap(m)
	if k != "c" {
		t.Error("Expected c")
	}
}
