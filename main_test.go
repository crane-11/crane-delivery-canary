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
	cases := []struct {
		name string
		a    string
		b    string
		want string
	}{
		// unequal inputs, one per argument position
		{"unequal/first-longer", "bbb", "a", "bbb"},
		{"unequal/second-longer", "a", "bbb", "bbb"},
		// equal-length but different: tie broken by returning the first argument
		{"equal/tie-keeps-first", "ab", "cd", "ab"},
		{"equal/tie-swapped-keeps-first", "cd", "ab", "cd"},
		// empty inputs
		{"empty/first-empty", "", "x", "x"},
		{"empty/second-empty", "x", "", "x"},
		{"empty/both-empty", "", "", ""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := longerString(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("longerString(%q, %q) = %q, want %q", tc.a, tc.b, got, tc.want)
			}
			// determinism: repeated calls with the same arguments agree
			if again := longerString(tc.a, tc.b); again != got {
				t.Fatalf("longerString(%q, %q) not deterministic: %q then %q", tc.a, tc.b, got, again)
			}
		})
	}
}
