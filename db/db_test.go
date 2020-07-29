package db

import "testing"

func TestConnect(t *testing.T) {
	want := "Hello, world."
	if got := Connect(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestGet(t *testing.T) {
	want := 1
	if got := Get("a"); got != want {
		t.Errorf("Get() = %q, want %q", got, want)
	}
}
