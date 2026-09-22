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

func TestLonger(t *testing.T) {
	// Unequal inputs: the longer string wins, regardless of argument order.
	if got := longer("a", "bbb"); got != "bbb" {
		t.Errorf("longer(%q, %q) = %q, want %q", "a", "bbb", got, "bbb")
	}
	if got := longer("bbb", "a"); got != "bbb" {
		t.Errorf("longer(%q, %q) = %q, want %q", "bbb", "a", got, "bbb")
	}
	if got := longer("abcdef", "ab"); got != "abcdef" {
		t.Errorf("longer(%q, %q) = %q, want %q", "abcdef", "ab", got, "abcdef")
	}

	// Equal-length, different-content inputs: the tie-break returns the FIRST argument.
	if got := longer("xy", "ab"); got != "xy" {
		t.Errorf("longer(%q, %q) = %q, want first argument %q", "xy", "ab", got, "xy")
	}
	if got := longer("xy", "ab"); got == "ab" {
		t.Errorf("longer(%q, %q) = %q, must not return the second argument", "xy", "ab", got)
	}
	if got := longer("abcd", "wxyz"); got != "abcd" {
		t.Errorf("longer(%q, %q) = %q, want first argument %q", "abcd", "wxyz", got, "abcd")
	}

	// Identical inputs.
	if got := longer("same", "same"); got != "same" {
		t.Errorf("longer(%q, %q) = %q, want %q", "same", "same", got, "same")
	}

	// Empty inputs: both empty, and each side empty against a non-empty string.
	if got := longer("", ""); got != "" {
		t.Errorf("longer(%q, %q) = %q, want %q", "", "", got, "")
	}
	if got := longer("", "x"); got != "x" {
		t.Errorf("longer(%q, %q) = %q, want %q", "", "x", got, "x")
	}
	if got := longer("x", ""); got != "x" {
		t.Errorf("longer(%q, %q) = %q, want %q", "x", "", got, "x")
	}
}
