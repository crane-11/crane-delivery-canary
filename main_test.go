package main

import "testing"

func TestOK(t *testing.T) {}

func TestAccept(t *testing.T) {
	if got := statusMessage(); got != "canary: unfiltered [canary-10]" {
		t.Fatalf("unexpected status message: %q", got)
	}
}
