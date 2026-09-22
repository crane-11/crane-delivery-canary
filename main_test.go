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
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{"unequal, longer first", "hello", "hi", "hello"},
		{"unequal, longer second", "hi", "hello", "hello"},
		{"equal length, first returned", "abc", "xyz", "abc"},
		{"equal length, single chars", "b", "a", "b"},
		{"empty vs non-empty", "", "x", "x"},
		{"non-empty vs empty", "x", "", "x"},
		{"empty vs empty", "", "", ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := longer(tc.a, tc.b); got != tc.want {
				t.Fatalf("longer(%q, %q) = %q, want %q", tc.a, tc.b, got, tc.want)
			}
		})
	}
	if got1, got2 := longer("hello", "hi"), longer("hello", "hi"); got1 != got2 || got1 != "hello" {
		t.Fatalf("longer is not deterministic: %q then %q", got1, got2)
	}
	if longer("", "") != "" {
		t.Fatalf("longer(\"\", \"\") not empty")
	}
	before := StatusDrillLabel() + "|" + BaseLabel() + "|" + statusMessage() + "|" + DrillOverlayProbe + "|" + DrillVersionMarker + "|" + DrillSweepMarker
	_ = longer("canary: unfiltered [canary-10b]", "drill-scenario-a")
	_ = longer("boundary", "probe")
	after := StatusDrillLabel() + "|" + BaseLabel() + "|" + statusMessage() + "|" + DrillOverlayProbe + "|" + DrillVersionMarker + "|" + DrillSweepMarker
	if before != after {
		t.Fatalf("longer altered package state: before %q, after %q", before, after)
	}
}
