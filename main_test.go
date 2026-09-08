package main

import (
	"strings"
	"testing"
)

func TestOK(t *testing.T) {}

func TestAccept(t *testing.T) {
	if statusMessage() == "" {
		t.Fatal("empty status message")
	}
	got := statusMessage()
	if got != "canary: unfiltered [canary-1]" {
		t.Fatalf("statusMessage() = %q, want %q", got, "canary: unfiltered [canary-1]")
	}
	const marker = " [canary-1]"
	if !strings.HasSuffix(got, marker) {
		t.Fatalf("statusMessage() = %q, want suffix %q", got, marker)
	}
	if strings.HasPrefix(got, marker) {
		t.Fatalf("statusMessage() = %q, marker must not be prepended", got)
	}
	if idx := strings.Index(got, marker); idx != len(got)-len(marker) {
		t.Fatalf("statusMessage() = %q, marker at index %d, want %d", got, idx, len(got)-len(marker))
	}
	if n := strings.Count(got, "[canary-1]"); n != 1 {
		t.Fatalf("statusMessage() = %q, marker occurs %d times, want exactly 1", got, n)
	}
	if base := strings.TrimSuffix(got, marker); base != "canary: unfiltered" {
		t.Fatalf("statusMessage() base = %q, want %q", base, "canary: unfiltered")
	}
	for i := 0; i < 3; i++ {
		if again := statusMessage(); again != got {
			t.Fatalf("statusMessage() call %d = %q, want byte-identical %q", i, again, got)
		}
	}
}
