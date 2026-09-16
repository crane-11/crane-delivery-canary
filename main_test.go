package main

import "testing"

func TestOK(t *testing.T) {}

func TestAccept(t *testing.T) {
	if got := statusMessage(); got != "canary: unfiltered [DRILL-BASE]" {
		t.Fatalf("statusMessage() = %q, want %q", got, "canary: unfiltered [DRILL-BASE]")
	}
}
