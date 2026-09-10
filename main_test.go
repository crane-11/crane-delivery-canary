package main

import "testing"

func TestOK(t *testing.T) {}

func TestAccept(t *testing.T) {
	if got := statusMessage(); got != "canary: ok" {
		t.Fatalf("statusMessage() = %q, want %q", got, "canary: ok")
	}
}
