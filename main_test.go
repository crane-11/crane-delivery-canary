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
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		// (a) unequal lengths: the longer string wins regardless of order.
		{name: "first longer", a: "abc", b: "de", want: "abc"},
		{name: "second longer", a: "de", b: "abc", want: "abc"},
		{name: "second clearly longer", a: "a", b: "abcd", want: "abcd"},
		// (b) equal length, distinct contents: the first argument wins on the tie.
		{name: "tie first wins", a: "abc", b: "xyz", want: "abc"},
		{name: "tie swapped first wins", a: "xyz", b: "abc", want: "xyz"},
		// (c) empty/edge inputs: length-comparison and tie boundaries.
		{name: "both empty", a: "", b: "", want: ""},
		{name: "first empty", a: "", b: "x", want: "x"},
		{name: "second empty", a: "x", b: "", want: "x"},
		{name: "first empty longer second", a: "", b: "ab", want: "ab"},
	}

	for _, tc := range tests {
		beforeA, beforeB := tc.a, tc.b
		gotA := longerString(tc.a, tc.b)
		if gotA != tc.want {
			t.Errorf("longerString(%q, %q) = %q, want %q", tc.a, tc.b, gotA, tc.want)
		}
		gotB := longerString(tc.a, tc.b)
		if gotB != gotA {
			t.Errorf("longerString(%q, %q) not pure: first call = %q, second call = %q", tc.a, tc.b, gotA, gotB)
		}
		if tc.a != beforeA || tc.b != beforeB {
			t.Errorf("longerString(%q, %q) mutated its arguments: got (%q, %q)", beforeA, beforeB, tc.a, tc.b)
		}
	}
}
