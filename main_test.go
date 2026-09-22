package main

import "testing"

func TestOK(t *testing.T) {}

func TestAccept(t *testing.T) {
	if got := statusMessage(); got != "canary: unfiltered [canary-10b]" {
		t.Fatalf("statusMessage() = %q, want %q", got, "canary: unfiltered [canary-10b]")
	}
	if got := StatusDrillLabel(); got != "canary: unfiltered [canary-10b] drill-scenario-a" {
		t.Fatalf("StatusDrillLabel() = %q, want %q", got, "canary: unfiltered [canary-10b] drill-scenario-a")
	}
	if got := BaseLabel(); got != "canary: unfiltered [drill-boundary]" {
		t.Fatalf("BaseLabel() = %q, want %q", got, "canary: unfiltered [drill-boundary]")
	}
	if BaseLabel() == StatusDrillLabel() {
		t.Fatalf("BaseLabel() == StatusDrillLabel() = %q, want distinct boundary", BaseLabel())
	}
	for i := 0; i+len(StatusDrill) <= len(BaseLabel()); i++ {
		if BaseLabel()[i:i+len(StatusDrill)] == StatusDrill {
			t.Fatalf("BaseLabel() = %q contains drill marker %q", BaseLabel(), StatusDrill)
		}
	}
}

func TestDrillVersionLabel(t *testing.T) {
	if got := DrillVersionLabel(); got != "version: canary-drill-marker-v1" {
		t.Fatalf("DrillVersionLabel() = %q, want %q", got, "version: canary-drill-marker-v1")
	}
}

func TestDrillSweepLabel(t *testing.T) {
	if got := DrillSweepLabel(); got != "sweep: canary-sweep-marker-v2" {
		t.Fatalf("DrillSweepLabel() = %q, want %q", got, "sweep: canary-sweep-marker-v2")
	}
}

func TestLongerString(t *testing.T) {
	// Equal inputs: equal length, so the first argument wins, even when the contents differ.
	if got := longerString("abc", "xyz"); got != "abc" {
		t.Fatalf("longerString(%q, %q) = %q, want %q", "abc", "xyz", got, "abc")
	}
	if got := longerString("same", "same"); got != "same" {
		t.Fatalf("longerString(%q, %q) = %q, want %q", "same", "same", got, "same")
	}
	// Unequal inputs, both argument orders: the longer one is always returned.
	if got := longerString("abcd", "ab"); got != "abcd" {
		t.Fatalf("longerString(%q, %q) = %q, want %q", "abcd", "ab", got, "abcd")
	}
	if got := longerString("ab", "abcd"); got != "abcd" {
		t.Fatalf("longerString(%q, %q) = %q, want %q", "ab", "abcd", got, "abcd")
	}
	// Empty inputs: an empty string paired with a non-empty one yields the non-empty string;
	// two empty strings yield the first (the empty string).
	if got := longerString("", "x"); got != "x" {
		t.Fatalf("longerString(%q, %q) = %q, want %q", "", "x", got, "x")
	}
	if got := longerString("x", ""); got != "x" {
		t.Fatalf("longerString(%q, %q) = %q, want %q", "x", "", got, "x")
	}
	if got := longerString("", ""); got != "" {
		t.Fatalf("longerString(%q, %q) = %q, want %q", "", "", got, "")
	}
}
