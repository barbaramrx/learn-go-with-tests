package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Babs")
	want := "Hello, Babs"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
