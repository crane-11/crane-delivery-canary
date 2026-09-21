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

// TestIsBlank verifies the pure helper IsBlank. The table covers the three
// required input categories: the empty string, whitespace-only inputs, and
// normal non-whitespace inputs. Each case exercises the helper through a
// subtest and asserts determinism over repeated calls.
//
// To keep this file transport-safe (no backslash escapes), the whitespace
// inputs are built from the whitespace rune literals 0x09 (tab), 0x0A
// (newline), 0x0D (carriage return) and 0x20 (space). These conversions
// produce exactly the whitespace characters IsBlank must classify as blank.
func TestIsBlank(t *testing.T) {
	const (
		space   = rune(0x20)
		tab     = rune(0x09)
		newline = rune(0x0A)
		cr      = rune(0x0D)
		nbsp    = rune(0xA0)
	)
	// The whitespace inputs below are built with string(rune) conversions so
	// this file contains no backslash escape sequences; each variable holds
	// exactly the whitespace characters its subtest name describes.
	wsSpace := string(space)
	wsTab := string(tab)
	wsNewline := string(newline)
	wsCRLF := string(cr) + string(newline)
	wsNbsp := string(nbsp)
	wsMixed := "  " + string(tab) + " " + string(newline) + " "

	cases := []struct {
		name string
		in   string
		want bool
	}{
		// (a) the empty string
		{"empty", "", true},
		// (b) whitespace-only inputs
		{"whitespace_only_space", wsSpace, true},
		{"whitespace_only_tab", wsTab, true},
		{"whitespace_only_newline", wsNewline, true},
		{"whitespace_only_crlf", wsCRLF, true},
		{"whitespace_only_nbsp", wsNbsp, true},
		{"whitespace_only_mixed", wsMixed, true},
		// (c) normal non-whitespace inputs
		{"normal_hello", "hello", false},
		{"normal_a", "a", false},
		{"normal_padded", " hello ", false},
		{"normal_tab_inside", "x" + string(tab) + "y", false},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := IsBlank(tc.in)
			if got != tc.want {
				t.Errorf("IsBlank(%q) = %v, want %v", tc.in, got, tc.want)
			}
			if again := IsBlank(tc.in); again != got {
				t.Errorf("IsBlank(%q) repeated = %v, first call gave %v, want deterministic", tc.in, got, again)
			}
		})
	}
}
