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

func TestIsBlank(t *testing.T) {
	// Empty input must report blank (true).
	if !isBlank("") {
		t.Fatalf("isBlank(%q) = false, want true", "")
	}

	// Whitespace-only input must report blank (true).
	for _, in := range []string{" ", "\t", "\n", "\r\n", "  \t \n "} {
		if !isBlank(in) {
			t.Fatalf("isBlank(%q) = false, want true", in)
		}
	}

	// Mixed-whitespace-only input must report blank (true).
	if !isBlank("\t\n\r \v\f") {
		t.Fatalf("isBlank(%q) = false, want true", "\t\n\r \v\f")
	}

	// Normal input must report not blank (false).
	for _, in := range []string{"hello", " hello ", "canary", "0"} {
		if isBlank(in) {
			t.Fatalf("isBlank(%q) = true, want false", in)
		}
	}
}
