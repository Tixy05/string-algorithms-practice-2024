package main

import "testing"

func TestExample(t *testing.T) {
	want := "TEST"
	if res := TestText(); res != want {
		t.Fatalf("expected: %s, got: %s", want, res)
	}
}
