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

// TestLongerString asserts LongerString's behaviour on equal, unequal, and
// empty inputs.
func TestLongerString(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		want string
	}{
		// (a) unequal-length inputs, both argument orderings.
		{"longer first", "abc", "de", "abc"},
		{"longer second", "de", "abc", "abc"},

		// (b) equal-length ties: want is always the first argument, so a
		// swapped tie rule fails. Left and right differ in content.
		{"tie returns first", "ab", "cd", "ab"},
		{"tie returns first reversed", "cd", "ab", "cd"},
		{"same-length non-empty tie", "zw", "xy", "zw"},

		// (c) empty inputs.
		{"both empty", "", "", ""},
		{"first empty", "", "x", "x"},
		{"second empty", "x", "", "x"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := LongerString(tc.a, tc.b); got != tc.want {
				t.Fatalf("LongerString(%q, %q) = %q, want %q", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
