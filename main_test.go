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
	// Unequal-length inputs: both ways must return the longer argument.
	if got, want := longer("ab", "abc"), "abc"; got != want {
		t.Errorf("longer(%q, %q) = %q, want %q", "ab", "abc", got, want)
	}
	if got, want := longer("abc", "ab"), "abc"; got != want {
		t.Errorf("longer(%q, %q) = %q, want %q", "abc", "ab", got, want)
	}

	// Equal-length distinct inputs: the tie must be broken by the first argument, both ways.
	if got, want := longer("xy", "ab"), "xy"; got != want {
		t.Errorf("longer(%q, %q) = %q, want %q", "xy", "ab", got, want)
	}
	if got, want := longer("ab", "xy"), "ab"; got != want {
		t.Errorf("longer(%q, %q) = %q, want %q", "ab", "xy", got, want)
	}
	if got, want := longer("zz", "aa"), "zz"; got != want {
		t.Errorf("longer(%q, %q) = %q, want %q", "zz", "aa", got, want)
	}

	// Empty inputs: including the empty/empty tie.
	if got, want := longer("", ""), ""; got != want {
		t.Errorf("longer(%q, %q) = %q, want %q", "", "", got, want)
	}
	if got, want := longer("", "a"), "a"; got != want {
		t.Errorf("longer(%q, %q) = %q, want %q", "", "a", got, want)
	}
	if got, want := longer("a", ""), "a"; got != want {
		t.Errorf("longer(%q, %q) = %q, want %q", "a", "", got, want)
	}

	// Determinism/purity: repeated calls with the same arguments return the same value.
	first := longer("ab", "abc")
	second := longer("ab", "abc")
	if first != "abc" || second != "abc" || first != second {
		t.Fatalf("longer(%q, %q) not deterministic: %q then %q, want %q", "ab", "abc", first, second, "abc")
	}
}
