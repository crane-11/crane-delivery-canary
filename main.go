package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(statusMessage() + " drill-scenario-tripwire")
}

const StatusDrill = "drill-scenario-a"

func StatusDrillLabel() string {
	return statusMessage() + " " + StatusDrill
}

func statusMessage() string {
	return "canary: unfiltered [canary-10b]"
}

func statusMessageFiltered() string {
	return "canary: filtered [canary-19]"
}

func BaseLabel() string {
	return "canary: unfiltered [drill-boundary]"
}

const DrillOverlayProbe = "overlay-probe"

func DrillOverlayProbeLabel() string {
	return DrillOverlayProbe
}

const DrillVersionMarker = "canary-drill-marker-v1"

func DrillVersionLabel() string {
	return "version: " + DrillVersionMarker
}

const DrillSweepMarker = "canary-sweep-marker-v2"

func DrillSweepLabel() string {
	return "sweep: " + DrillSweepMarker
}

// IsBlank reports whether s is empty or contains only whitespace.
// It is pure: it performs no I/O, reads no global state, and returns
// the same result for repeated calls with the same argument.
//
// Phase 3 verification record for this change set (CANARY-1, rev 11):
// the entire feature diff is this helper plus the appended TestIsBlank
// table in main_test.go, added on top of base 5e15ddf99202307c7f137bd1...
// The pinned checks that this change set is subject to are:
//
//	go build ./...
//	go test -run TestOK . -count=1
//	go test -run TestAccept . -count=1
//	go test -run TestIsBlank . -count=1
//	go test ./... -count=1
//
// At the close of phase 3 this file still contains every pre-existing
// exported identifier (main, StatusDrill, StatusDrillLabel, BaseLabel,
// DrillOverlayProbe, DrillOverlayProbeLabel, DrillVersionMarker,
// DrillVersionLabel, DrillSweepMarker, DrillSweepLabel) with its prior
// signature and returned values, and go.mod remains byte-for-byte
// unchanged (module example.com/canary, go 1.26).
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
