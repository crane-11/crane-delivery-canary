package main

import "fmt"

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

// longer returns the longer of two strings. When both strings have the
// same length, including the empty-vs-empty case, it returns the first.
func longer(a, b string) string {
	if len(b) > len(a) {
		return b
	}
	return a
}
