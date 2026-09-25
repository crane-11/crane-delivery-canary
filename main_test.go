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
	// unequal lengths, either argument order returns the longer
	if got := Longer("a", "bb"); got != "bb" {
		t.Fatalf("Longer(\"a\",\"bb\") = %q, want %q", got, "bb")
	}
	if got := Longer("bb", "a"); got != "bb" {
		t.Fatalf("Longer(\"bb\",\"a\") = %q, want %q", got, "bb")
	}
	// equal lengths return the first argument (tie-break)
	if got := Longer("xy", "ab"); got != "xy" {
		t.Fatalf("Longer(\"xy\",\"ab\") = %q, want %q", got, "xy")
	}
	if got := Longer("zz", "ab"); got != "zz" {
		t.Fatalf("Longer(\"zz\",\"ab\") = %q, want %q", got, "zz")
	}
	// empty inputs
	if got := Longer("", ""); got != "" {
		t.Fatalf("Longer(\"\",\"\") = %q, want %q", got, "")
	}
	if got := Longer("", "a"); got != "a" {
		t.Fatalf("Longer(\"\",\"a\") = %q, want %q", got, "a")
	}
	if got := Longer("a", ""); got != "a" {
		t.Fatalf("Longer(\"a\",\"\") = %q, want %q", got, "a")
	}
}
