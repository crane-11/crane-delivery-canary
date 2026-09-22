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

// TestLongerString asserts LongerString returns the longer operand, the first
// on equal-length ties, and behaves sanely on empty inputs.
func TestLongerString(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		// Unequal-length inputs: the longer string wins (AC1).
		{"longer first", "abc", "de", "abc"},
		{"longer second", "de", "abc", "abc"},
		// Equal-length distinct inputs: the FIRST argument wins the tie (AC2).
		{"tie first wins a", "ab", "cd", "ab"},
		{"tie first wins b", "cd", "ab", "cd"},
		// Byte semantics: len counts bytes, not runes.
		{"multi-byte rune", "é", "ab", "ab"},
		// Empty inputs.
		{"both empty", "", "", ""},
		{"empty first", "", "abc", "abc"},
		{"empty second", "abc", "", "abc"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongerString(tc.a, tc.b); got != tc.want {
				t.Errorf("LongerString(%q, %q) = %q, want %q", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
