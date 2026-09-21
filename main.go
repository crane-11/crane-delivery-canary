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

// isBlank reports whether s is empty or consists only of whitespace.
// It is a pure helper: it depends on no globals, mutates no state, and
// uses only the Go standard library.
func isBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
