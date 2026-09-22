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

// TestLonger locks longer's behaviour across the three required input
// categories — unequal lengths, equal-length ties (first argument wins) and
// empty inputs — and runs alongside the untouched TestOK/TestAccept set.
func TestLonger(t *testing.T) {
	cases := []struct {
		name string
		a, b string
		want string
	}{
		{name: "unequal, longer second", a: "ab", b: "abcd", want: "abcd"},
		{name: "unequal, longer first", a: "abcd", b: "ab", want: "abcd"},
		{name: "equal tie returns the first", a: "abc", b: "xyz", want: "abc"},
		{name: "equal identical tie returns the first", a: "xy", b: "xy", want: "xy"},
		{name: "empty versus non-empty", a: "", b: "x", want: "x"},
		{name: "non-empty versus empty", a: "x", b: "", want: "x"},
		{name: "both empty", a: "", b: "", want: ""},
	}
	for _, tc := range cases {
		if got := longer(tc.a, tc.b); got != tc.want {
			t.Errorf("%s: longer(%q, %q) = %q, want %q", tc.name, tc.a, tc.b, got, tc.want)
		}
	}
}
