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
	if got := IsBlank(""); !got {
		t.Fatalf(`IsBlank("") = %v, want true`, got)
	}
	if got := IsBlank("   \t\n"); !got {
		t.Fatalf(`IsBlank("   \t\n") = %v, want true`, got)
	}
	if got := IsBlank("canary"); got {
		t.Fatalf(`IsBlank("canary") = %v, want false`, got)
	}
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"empty input", "", true},
		{"single space", " ", true},
		{"multiple spaces", "   ", true},
		{"tab only", "\t", true},
		{"newline only", "\n", true},
		{"spaces tab and newline", " \t\n ", true},
		{"spaces tab and newline trailing", "   \t\n", true},
		{"tab and newline", "\t\n", true},
		{"normal input", "canary", false},
		{"normal input with surrounding whitespace", " a ", false},
		{"normal multiword input", "canary drill", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsBlank(tc.in)
			want := tc.want
			if got != want {
				t.Errorf("IsBlank(%q) = %v, want %v", tc.in, got, want)
				return
			}
		})
	}
}
